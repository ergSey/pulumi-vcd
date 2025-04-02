# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs

allowApiTokenFile: Optional[bool]
"""
Set this to true if you understand the security risks of using API token files and would like to suppress the warnings
"""

allowServiceAccountTokenFile: Optional[bool]
"""
Set this to true if you understand the security risks of using Service Account token files and would like to suppress
the warnings
"""

allowUnverifiedSsl: Optional[bool]
"""
If set, VCDClient will permit unverifiable SSL certificates.
"""

apiToken: Optional[str]
"""
The API token used instead of username/password for VCD API operations. (Requires VCD 10.3.1+)
"""

apiTokenFile: Optional[str]
"""
The API token file instead of username/password for VCD API operations. (Requires VCD 10.3.1+)
"""

authType: Optional[str]
"""
'integrated', 'saml_adfs', 'token', 'api_token', 'api_token_file' and 'service_account_token_file' are supported.
'integrated' is default.
"""

ignoreMetadataChanges: Optional[str]
"""
Defines a set of `metadata_entry` that need to be ignored by this provider. All filters on this attribute are computed
with a logical AND
"""

importSeparator: Optional[str]

logging: Optional[bool]
"""
If set, it will enable logging of API requests and responses
"""

loggingFile: Optional[str]
"""
Defines the full name of the logging file for API calls (requires 'logging')
"""

maxRetryTimeout: Optional[int]
"""
Max num seconds to wait for successful response when operating on resources within vCloud (defaults to 60)
"""

org: Optional[str]
"""
The VCD Org for API operations
"""

password: Optional[str]
"""
The user password for VCD API operations.
"""

samlAdfsCookie: Optional[str]
"""
Allows to specify custom cookie for ADFS server lookup. '{{.Org}}' is replaced by real Org - e.g. 'sso-preferred=yes;
sso_redirect_org={{.Org}}'
"""

samlAdfsRptId: Optional[str]
"""
Allows to specify custom Relaying Party Trust Identifier for auth_type=saml_adfs
"""

serviceAccountTokenFile: Optional[str]
"""
The Service Account API token file instead of username/password for VCD API operations. (Requires VCD 10.4.0+)
"""

sysorg: Optional[str]
"""
The VCD Org for user authentication
"""

token: Optional[str]
"""
The token used instead of username/password for VCD API operations.
"""

url: Optional[str]
"""
The VCD url for VCD API operations.
"""

user: Optional[str]
"""
The user name for VCD API operations.
"""

vdc: Optional[str]
"""
The VDC for API operations
"""

