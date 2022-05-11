// Code generated by "stringer -type=MemoryKind -trimprefix=MemoryKind"; DO NOT EDIT.

package common

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MemoryKindUnknown-0]
	_ = x[MemoryKindBool-1]
	_ = x[MemoryKindAddress-2]
	_ = x[MemoryKindString-3]
	_ = x[MemoryKindCharacter-4]
	_ = x[MemoryKindMetaType-5]
	_ = x[MemoryKindNumber-6]
	_ = x[MemoryKindArrayBase-7]
	_ = x[MemoryKindDictionaryBase-8]
	_ = x[MemoryKindCompositeBase-9]
	_ = x[MemoryKindSimpleCompositeBase-10]
	_ = x[MemoryKindOptional-11]
	_ = x[MemoryKindNil-12]
	_ = x[MemoryKindVoid-13]
	_ = x[MemoryKindTypeValue-14]
	_ = x[MemoryKindPathValue-15]
	_ = x[MemoryKindCapabilityValue-16]
	_ = x[MemoryKindLinkValue-17]
	_ = x[MemoryKindStorageReferenceValue-18]
	_ = x[MemoryKindEphemeralReferenceValue-19]
	_ = x[MemoryKindInterpretedFunction-20]
	_ = x[MemoryKindHostFunction-21]
	_ = x[MemoryKindBoundFunction-22]
	_ = x[MemoryKindBigInt-23]
	_ = x[MemoryKindSimpleComposite-24]
	_ = x[MemoryKindAtreeArrayDataSlab-25]
	_ = x[MemoryKindAtreeArrayMetaDataSlab-26]
	_ = x[MemoryKindAtreeArrayElementOverhead-27]
	_ = x[MemoryKindAtreeMapDataSlab-28]
	_ = x[MemoryKindAtreeMapMetaDataSlab-29]
	_ = x[MemoryKindAtreeMapElementOverhead-30]
	_ = x[MemoryKindAtreeMapPreAllocatedElement-31]
	_ = x[MemoryKindPrimitiveStaticType-32]
	_ = x[MemoryKindCompositeStaticType-33]
	_ = x[MemoryKindInterfaceStaticType-34]
	_ = x[MemoryKindVariableSizedStaticType-35]
	_ = x[MemoryKindConstantSizedStaticType-36]
	_ = x[MemoryKindDictionaryStaticType-37]
	_ = x[MemoryKindOptionalStaticType-38]
	_ = x[MemoryKindRestrictedStaticType-39]
	_ = x[MemoryKindReferenceStaticType-40]
	_ = x[MemoryKindCapabilityStaticType-41]
	_ = x[MemoryKindFunctionStaticType-42]
	_ = x[MemoryKindCadenceVoid-43]
	_ = x[MemoryKindCadenceOptional-44]
	_ = x[MemoryKindCadenceBool-45]
	_ = x[MemoryKindCadenceString-46]
	_ = x[MemoryKindCadenceCharacter-47]
	_ = x[MemoryKindCadenceAddress-48]
	_ = x[MemoryKindCadenceInt-49]
	_ = x[MemoryKindCadenceNumber-50]
	_ = x[MemoryKindCadenceArrayBase-51]
	_ = x[MemoryKindCadenceArrayLength-52]
	_ = x[MemoryKindCadenceDictionaryBase-53]
	_ = x[MemoryKindCadenceDictionarySize-54]
	_ = x[MemoryKindCadenceKeyValuePair-55]
	_ = x[MemoryKindCadenceStructBase-56]
	_ = x[MemoryKindCadenceStructSize-57]
	_ = x[MemoryKindCadenceResourceBase-58]
	_ = x[MemoryKindCadenceResourceSize-59]
	_ = x[MemoryKindCadenceEventBase-60]
	_ = x[MemoryKindCadenceEventSize-61]
	_ = x[MemoryKindCadenceContractBase-62]
	_ = x[MemoryKindCadenceContractSize-63]
	_ = x[MemoryKindCadenceEnumBase-64]
	_ = x[MemoryKindCadenceEnumSize-65]
	_ = x[MemoryKindCadenceLink-66]
	_ = x[MemoryKindCadencePath-67]
	_ = x[MemoryKindCadenceTypeValue-68]
	_ = x[MemoryKindCadenceCapability-69]
	_ = x[MemoryKindCadenceSimpleType-70]
	_ = x[MemoryKindCadenceOptionalType-71]
	_ = x[MemoryKindCadenceVariableSizedArrayType-72]
	_ = x[MemoryKindCadenceConstantSizedArrayType-73]
	_ = x[MemoryKindCadenceDictionaryType-74]
	_ = x[MemoryKindCadenceField-75]
	_ = x[MemoryKindCadenceParameter-76]
	_ = x[MemoryKindCadenceStructType-77]
	_ = x[MemoryKindCadenceResourceType-78]
	_ = x[MemoryKindCadenceEventType-79]
	_ = x[MemoryKindCadenceContractType-80]
	_ = x[MemoryKindCadenceStructInterfaceType-81]
	_ = x[MemoryKindCadenceResourceInterfaceType-82]
	_ = x[MemoryKindCadenceContractInterfaceType-83]
	_ = x[MemoryKindCadenceFunctionType-84]
	_ = x[MemoryKindCadenceReferenceType-85]
	_ = x[MemoryKindCadenceRestrictedType-86]
	_ = x[MemoryKindCadenceCapabilityType-87]
	_ = x[MemoryKindCadenceEnumType-88]
	_ = x[MemoryKindRawString-89]
	_ = x[MemoryKindAddressLocation-90]
	_ = x[MemoryKindBytes-91]
	_ = x[MemoryKindVariable-92]
	_ = x[MemoryKindCompositeTypeInfo-93]
	_ = x[MemoryKindCompositeField-94]
	_ = x[MemoryKindInvocation-95]
	_ = x[MemoryKindValueToken-96]
	_ = x[MemoryKindSyntaxToken-97]
	_ = x[MemoryKindSpaceToken-98]
	_ = x[MemoryKindProgram-99]
	_ = x[MemoryKindIdentifier-100]
	_ = x[MemoryKindArgument-101]
	_ = x[MemoryKindBlock-102]
	_ = x[MemoryKindFunctionBlock-103]
	_ = x[MemoryKindParameter-104]
	_ = x[MemoryKindParameterList-105]
	_ = x[MemoryKindTransfer-106]
	_ = x[MemoryKindMembers-107]
	_ = x[MemoryKindTypeAnnotation-108]
	_ = x[MemoryKindDictionaryEntry-109]
	_ = x[MemoryKindFunctionDeclaration-110]
	_ = x[MemoryKindCompositeDeclaration-111]
	_ = x[MemoryKindInterfaceDeclaration-112]
	_ = x[MemoryKindEnumCaseDeclaration-113]
	_ = x[MemoryKindFieldDeclaration-114]
	_ = x[MemoryKindTransactionDeclaration-115]
	_ = x[MemoryKindImportDeclaration-116]
	_ = x[MemoryKindVariableDeclaration-117]
	_ = x[MemoryKindSpecialFunctionDeclaration-118]
	_ = x[MemoryKindPragmaDeclaration-119]
	_ = x[MemoryKindAssignmentStatement-120]
	_ = x[MemoryKindBreakStatement-121]
	_ = x[MemoryKindContinueStatement-122]
	_ = x[MemoryKindEmitStatement-123]
	_ = x[MemoryKindExpressionStatement-124]
	_ = x[MemoryKindForStatement-125]
	_ = x[MemoryKindIfStatement-126]
	_ = x[MemoryKindReturnStatement-127]
	_ = x[MemoryKindSwapStatement-128]
	_ = x[MemoryKindSwitchStatement-129]
	_ = x[MemoryKindWhileStatement-130]
	_ = x[MemoryKindBooleanExpression-131]
	_ = x[MemoryKindNilExpression-132]
	_ = x[MemoryKindStringExpression-133]
	_ = x[MemoryKindIntegerExpression-134]
	_ = x[MemoryKindFixedPointExpression-135]
	_ = x[MemoryKindArrayExpression-136]
	_ = x[MemoryKindDictionaryExpression-137]
	_ = x[MemoryKindIdentifierExpression-138]
	_ = x[MemoryKindInvocationExpression-139]
	_ = x[MemoryKindMemberExpression-140]
	_ = x[MemoryKindIndexExpression-141]
	_ = x[MemoryKindConditionalExpression-142]
	_ = x[MemoryKindUnaryExpression-143]
	_ = x[MemoryKindBinaryExpression-144]
	_ = x[MemoryKindFunctionExpression-145]
	_ = x[MemoryKindCastingExpression-146]
	_ = x[MemoryKindCreateExpression-147]
	_ = x[MemoryKindDestroyExpression-148]
	_ = x[MemoryKindReferenceExpression-149]
	_ = x[MemoryKindForceExpression-150]
	_ = x[MemoryKindPathExpression-151]
	_ = x[MemoryKindConstantSizedType-152]
	_ = x[MemoryKindDictionaryType-153]
	_ = x[MemoryKindFunctionType-154]
	_ = x[MemoryKindInstantiationType-155]
	_ = x[MemoryKindNominalType-156]
	_ = x[MemoryKindOptionalType-157]
	_ = x[MemoryKindReferenceType-158]
	_ = x[MemoryKindRestrictedType-159]
	_ = x[MemoryKindVariableSizedType-160]
	_ = x[MemoryKindPosition-161]
	_ = x[MemoryKindRange-162]
	_ = x[MemoryKindElaboration-163]
	_ = x[MemoryKindActivation-164]
	_ = x[MemoryKindActivationEntries-165]
	_ = x[MemoryKindVariableSizedSemaType-166]
	_ = x[MemoryKindConstantSizedSemaType-167]
	_ = x[MemoryKindDictionarySemaType-168]
	_ = x[MemoryKindOptionalSemaType-169]
	_ = x[MemoryKindRestrictedSemaType-170]
	_ = x[MemoryKindReferenceSemaType-171]
	_ = x[MemoryKindCapabilitySemaType-172]
	_ = x[MemoryKindOrderedMap-173]
	_ = x[MemoryKindOrderedMapEntryList-174]
	_ = x[MemoryKindOrderedMapEntry-175]
	_ = x[MemoryKindLast-176]
}

const _MemoryKind_name = "UnknownBoolAddressStringCharacterMetaTypeNumberArrayBaseDictionaryBaseCompositeBaseSimpleCompositeBaseOptionalNilVoidTypeValuePathValueCapabilityValueLinkValueStorageReferenceValueEphemeralReferenceValueInterpretedFunctionHostFunctionBoundFunctionBigIntSimpleCompositeAtreeArrayDataSlabAtreeArrayMetaDataSlabAtreeArrayElementOverheadAtreeMapDataSlabAtreeMapMetaDataSlabAtreeMapElementOverheadAtreeMapPreAllocatedElementPrimitiveStaticTypeCompositeStaticTypeInterfaceStaticTypeVariableSizedStaticTypeConstantSizedStaticTypeDictionaryStaticTypeOptionalStaticTypeRestrictedStaticTypeReferenceStaticTypeCapabilityStaticTypeFunctionStaticTypeCadenceVoidCadenceOptionalCadenceBoolCadenceStringCadenceCharacterCadenceAddressCadenceIntCadenceNumberCadenceArrayBaseCadenceArrayLengthCadenceDictionaryBaseCadenceDictionarySizeCadenceKeyValuePairCadenceStructBaseCadenceStructSizeCadenceResourceBaseCadenceResourceSizeCadenceEventBaseCadenceEventSizeCadenceContractBaseCadenceContractSizeCadenceEnumBaseCadenceEnumSizeCadenceLinkCadencePathCadenceTypeValueCadenceCapabilityCadenceSimpleTypeCadenceOptionalTypeCadenceVariableSizedArrayTypeCadenceConstantSizedArrayTypeCadenceDictionaryTypeCadenceFieldCadenceParameterCadenceStructTypeCadenceResourceTypeCadenceEventTypeCadenceContractTypeCadenceStructInterfaceTypeCadenceResourceInterfaceTypeCadenceContractInterfaceTypeCadenceFunctionTypeCadenceReferenceTypeCadenceRestrictedTypeCadenceCapabilityTypeCadenceEnumTypeRawStringAddressLocationBytesVariableCompositeTypeInfoCompositeFieldInvocationValueTokenSyntaxTokenSpaceTokenProgramIdentifierArgumentBlockFunctionBlockParameterParameterListTransferMembersTypeAnnotationDictionaryEntryFunctionDeclarationCompositeDeclarationInterfaceDeclarationEnumCaseDeclarationFieldDeclarationTransactionDeclarationImportDeclarationVariableDeclarationSpecialFunctionDeclarationPragmaDeclarationAssignmentStatementBreakStatementContinueStatementEmitStatementExpressionStatementForStatementIfStatementReturnStatementSwapStatementSwitchStatementWhileStatementBooleanExpressionNilExpressionStringExpressionIntegerExpressionFixedPointExpressionArrayExpressionDictionaryExpressionIdentifierExpressionInvocationExpressionMemberExpressionIndexExpressionConditionalExpressionUnaryExpressionBinaryExpressionFunctionExpressionCastingExpressionCreateExpressionDestroyExpressionReferenceExpressionForceExpressionPathExpressionConstantSizedTypeDictionaryTypeFunctionTypeInstantiationTypeNominalTypeOptionalTypeReferenceTypeRestrictedTypeVariableSizedTypePositionRangeElaborationActivationActivationEntriesVariableSizedSemaTypeConstantSizedSemaTypeDictionarySemaTypeOptionalSemaTypeRestrictedSemaTypeReferenceSemaTypeCapabilitySemaTypeOrderedMapOrderedMapEntryListOrderedMapEntryLast"

var _MemoryKind_index = [...]uint16{0, 7, 11, 18, 24, 33, 41, 47, 56, 70, 83, 102, 110, 113, 117, 126, 135, 150, 159, 180, 203, 222, 234, 247, 253, 268, 286, 308, 333, 349, 369, 392, 419, 438, 457, 476, 499, 522, 542, 560, 580, 599, 619, 637, 648, 663, 674, 687, 703, 717, 727, 740, 756, 774, 795, 816, 835, 852, 869, 888, 907, 923, 939, 958, 977, 992, 1007, 1018, 1029, 1045, 1062, 1079, 1098, 1127, 1156, 1177, 1189, 1205, 1222, 1241, 1257, 1276, 1302, 1330, 1358, 1377, 1397, 1418, 1439, 1454, 1463, 1478, 1483, 1491, 1508, 1522, 1532, 1542, 1553, 1563, 1570, 1580, 1588, 1593, 1606, 1615, 1628, 1636, 1643, 1657, 1672, 1691, 1711, 1731, 1750, 1766, 1788, 1805, 1824, 1850, 1867, 1886, 1900, 1917, 1930, 1949, 1961, 1972, 1987, 2000, 2015, 2029, 2046, 2059, 2075, 2092, 2112, 2127, 2147, 2167, 2187, 2203, 2218, 2239, 2254, 2270, 2288, 2305, 2321, 2338, 2357, 2372, 2386, 2403, 2417, 2429, 2446, 2457, 2469, 2482, 2496, 2513, 2521, 2526, 2537, 2547, 2564, 2585, 2606, 2624, 2640, 2658, 2675, 2693, 2703, 2722, 2737, 2741}

func (i MemoryKind) String() string {
	if i >= MemoryKind(len(_MemoryKind_index)-1) {
		return "MemoryKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MemoryKind_name[_MemoryKind_index[i]:_MemoryKind_index[i+1]]
}
