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

import types

__config__ = pulumi.Config('vcd')


class _ExportableConfig(types.ModuleType):
    @property
    def allow_api_token_file(self) -> Optional[bool]:
        """
        Set this to true if you understand the security risks of using API token files and would like to suppress the warnings
        """
        return __config__.get_bool('allowApiTokenFile')

    @property
    def allow_service_account_token_file(self) -> Optional[bool]:
        """
        Set this to true if you understand the security risks of using Service Account token files and would like to suppress
        the warnings
        """
        return __config__.get_bool('allowServiceAccountTokenFile')

    @property
    def allow_unverified_ssl(self) -> Optional[bool]:
        """
        If set, VCDClient will permit unverifiable SSL certificates.
        """
        return __config__.get_bool('allowUnverifiedSsl')

    @property
    def api_token(self) -> Optional[str]:
        """
        The API token used instead of username/password for VCD API operations. (Requires VCD 10.3.1+)
        """
        return __config__.get('apiToken')

    @property
    def api_token_file(self) -> Optional[str]:
        """
        The API token file instead of username/password for VCD API operations. (Requires VCD 10.3.1+)
        """
        return __config__.get('apiTokenFile')

    @property
    def auth_type(self) -> Optional[str]:
        """
        'integrated', 'saml_adfs', 'token', 'api_token', 'api_token_file' and 'service_account_token_file' are supported.
        'integrated' is default.
        """
        return __config__.get('authType')

    @property
    def ignore_metadata_changes(self) -> Optional[str]:
        """
        Defines a set of `metadata_entry` that need to be ignored by this provider. All filters on this attribute are computed
        with a logical AND
        """
        return __config__.get('ignoreMetadataChanges')

    @property
    def import_separator(self) -> Optional[str]:
        return __config__.get('importSeparator')

    @property
    def logging(self) -> Optional[bool]:
        """
        If set, it will enable logging of API requests and responses
        """
        return __config__.get_bool('logging')

    @property
    def logging_file(self) -> Optional[str]:
        """
        Defines the full name of the logging file for API calls (requires 'logging')
        """
        return __config__.get('loggingFile')

    @property
    def max_retry_timeout(self) -> Optional[int]:
        """
        Max num seconds to wait for successful response when operating on resources within vCloud (defaults to 60)
        """
        return __config__.get_int('maxRetryTimeout')

    @property
    def org(self) -> Optional[str]:
        """
        The VCD Org for API operations
        """
        return __config__.get('org')

    @property
    def password(self) -> Optional[str]:
        """
        The user password for VCD API operations.
        """
        return __config__.get('password')

    @property
    def saml_adfs_cookie(self) -> Optional[str]:
        """
        Allows to specify custom cookie for ADFS server lookup. '{{.Org}}' is replaced by real Org - e.g. 'sso-preferred=yes;
        sso_redirect_org={{.Org}}'
        """
        return __config__.get('samlAdfsCookie')

    @property
    def saml_adfs_rpt_id(self) -> Optional[str]:
        """
        Allows to specify custom Relaying Party Trust Identifier for auth_type=saml_adfs
        """
        return __config__.get('samlAdfsRptId')

    @property
    def service_account_token_file(self) -> Optional[str]:
        """
        The Service Account API token file instead of username/password for VCD API operations. (Requires VCD 10.4.0+)
        """
        return __config__.get('serviceAccountTokenFile')

    @property
    def sysorg(self) -> Optional[str]:
        """
        The VCD Org for user authentication
        """
        return __config__.get('sysorg')

    @property
    def token(self) -> Optional[str]:
        """
        The token used instead of username/password for VCD API operations.
        """
        return __config__.get('token')

    @property
    def url(self) -> Optional[str]:
        """
        The VCD url for VCD API operations.
        """
        return __config__.get('url')

    @property
    def user(self) -> Optional[str]:
        """
        The user name for VCD API operations.
        """
        return __config__.get('user')

    @property
    def vdc(self) -> Optional[str]:
        """
        The VDC for API operations
        """
        return __config__.get('vdc')

