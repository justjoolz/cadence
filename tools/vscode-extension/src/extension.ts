import {
  ExtensionContext,
  commands,
  workspace,
  window,
  ColorPresentation
} from "vscode";
import { LanguageClient } from "vscode-languageclient";

export function activate(ctx: ExtensionContext) {
  detectLaunchConfigurationChanges();

  function registerCommand(command: string, callback: (...args: any[]) => any) {
    ctx.subscriptions.push(commands.registerCommand(command, callback));
  }

  let client = startServer(ctx);

  registerCommand("cadence.restartServer", async () => {
    if (!client) {
      return;
    }
    await client.stop();
    client = startServer(ctx);
  });
}

function startServer(ctx: ExtensionContext): LanguageClient | undefined {
  const languageServerCommand: string | undefined = workspace
    .getConfiguration("cadence")
    .get("languageServerCommand");

  if (!languageServerCommand) {
    window.showWarningMessage("Missing command to start the Cadence language server");
    return;
  }

  const startLanguageServerCommandAndArgs = languageServerCommand.split(/\s+/)
  if (startLanguageServerCommandAndArgs.length < 1) {
    window.showWarningMessage("Malformed language server command")
    return
  }

  const command = startLanguageServerCommandAndArgs[0]
  const args = startLanguageServerCommandAndArgs.splice(1)

  const client = new LanguageClient(
    "cadence",
    "Cadence",
    {
      command,
      args,
    },
    {
      documentSelector: [{ scheme: "file", language: "cadence" }],
      synchronize: {
        configurationSection: "cadence"
      },
      initializationOptions: {
        "test": "test",
      },
    }
  );

  client
    .onReady()
    .then(() => {
      return window.showInformationMessage("Cadence language server started");
    })
    .catch(error => {
      return window.showErrorMessage(
        `Cadence language server failed to start: ${error}`
      );
    });

  let languageServerDisposable = client.start();
  ctx.subscriptions.push(languageServerDisposable);

  return client;
}

function detectLaunchConfigurationChanges() {
  workspace.onDidChangeConfiguration(e => {
    const promptRestartKeys = ["languageServerPath"];
    const shouldPromptRestart = promptRestartKeys.some(key =>
      e.affectsConfiguration(`cadence.${key}`)
    );
    if (shouldPromptRestart) {
      window
        .showInformationMessage(
          "Server launch configuration change detected. Reload the window for changes to take effect",
          "Reload Window",
          "Not now"
        )
        .then(choice => {
          if (choice === "Reload Window") {
            commands.executeCommand("workbench.action.reloadWindow");
          }
        });
    }
  });
}

export function deactivate() {}
