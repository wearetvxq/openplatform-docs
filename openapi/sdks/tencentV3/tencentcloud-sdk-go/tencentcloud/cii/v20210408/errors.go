// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20210408

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 企业认证失败。
	AUTHFAILURE_ENTERPRISE = "AuthFailure.Enterprise"

	// 个人认证失败。
	AUTHFAILURE_PERSONAL = "AuthFailure.Personal"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"
)
