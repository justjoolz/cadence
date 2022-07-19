
pub contract Test {

    pub struct Blockchain {

        pub let backend: AnyStruct{BlockchainBackend}

        init(backend: AnyStruct{BlockchainBackend}) {
            self.backend = backend
        }

        pub fun executeScript(_ script: String): ScriptResult {
            return self.backend.executeScript(script)
        }

        //pub fun addTransaction(_ transaction: Transaction) {
        //    self.backend.addTransaction(transaction)
        //}

        /// Executes the next transaction, if any.
        /// Returns the result of the transaction, or nil if no transaction was scheduled.
        ///
        //pub fun executeNextTransaction(): TransactionResult? {
        //    return self.backend.executeNextTransaction()
        //}

        //pub fun commitBlock() {
        //    self.backend.commitBlock()
        //}

        //pub fun executeTransaction(_ transaction: Transaction): TransactionResult {
        //    self.addTransaction(transaction)
        //    let result = self.executeNextTransaction()!
        //    self.commitBlock()
        //    return result
        //}

        //pub fun executeTransactions(_ transactions: [Transaction]): [TransactionResult] {
        //    for transaction in transactions {
        //        self.addTransaction(transaction)
        //    }

        //    var results: [TransactionResult] = []
        //    for transaction in transactions {
        //        let result = self.executeNextTransaction()!
        //        results.append(result)
        //    }

        //    self.commitBlock()
        //    return results
        //}
    }

    pub enum ResultStatus: UInt8 {
        pub case succeeded
        pub case failed
    }

    pub struct TransactionResult {
        pub let status: ResultStatus

        init(_ status: ResultStatus) {
            self.status = status
        }
    }

    pub struct ScriptResult {
        pub let status: ResultStatus
        pub let returnValue: AnyStruct?

        init(_ status: ResultStatus, _ returnValue: AnyStruct?) {
            self.status = status
            self.returnValue = returnValue
        }
    }

    pub struct interface BlockchainBackend {

        // fun addTransaction(_ transaction: Transaction)

        // fun executeNextTransaction(): TransactionResult?

        // fun commitBlock()

        pub fun executeScript(_ script: String): ScriptResult
    }
}
