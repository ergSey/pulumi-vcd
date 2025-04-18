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
from . import _utilities

__all__ = [
    'GetApiFilterResult',
    'AwaitableGetApiFilterResult',
    'get_api_filter',
    'get_api_filter_output',
]

@pulumi.output_type
class GetApiFilterResult:
    """
    A collection of values returned by getApiFilter.
    """
    def __init__(__self__, api_filter_id=None, external_endpoint_id=None, id=None, url_matcher_pattern=None, url_matcher_scope=None):
        if api_filter_id and not isinstance(api_filter_id, str):
            raise TypeError("Expected argument 'api_filter_id' to be a str")
        pulumi.set(__self__, "api_filter_id", api_filter_id)
        if external_endpoint_id and not isinstance(external_endpoint_id, str):
            raise TypeError("Expected argument 'external_endpoint_id' to be a str")
        pulumi.set(__self__, "external_endpoint_id", external_endpoint_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if url_matcher_pattern and not isinstance(url_matcher_pattern, str):
            raise TypeError("Expected argument 'url_matcher_pattern' to be a str")
        pulumi.set(__self__, "url_matcher_pattern", url_matcher_pattern)
        if url_matcher_scope and not isinstance(url_matcher_scope, str):
            raise TypeError("Expected argument 'url_matcher_scope' to be a str")
        pulumi.set(__self__, "url_matcher_scope", url_matcher_scope)

    @property
    @pulumi.getter(name="apiFilterId")
    def api_filter_id(self) -> str:
        return pulumi.get(self, "api_filter_id")

    @property
    @pulumi.getter(name="externalEndpointId")
    def external_endpoint_id(self) -> str:
        return pulumi.get(self, "external_endpoint_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="urlMatcherPattern")
    def url_matcher_pattern(self) -> str:
        return pulumi.get(self, "url_matcher_pattern")

    @property
    @pulumi.getter(name="urlMatcherScope")
    def url_matcher_scope(self) -> str:
        return pulumi.get(self, "url_matcher_scope")


class AwaitableGetApiFilterResult(GetApiFilterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApiFilterResult(
            api_filter_id=self.api_filter_id,
            external_endpoint_id=self.external_endpoint_id,
            id=self.id,
            url_matcher_pattern=self.url_matcher_pattern,
            url_matcher_scope=self.url_matcher_scope)


def get_api_filter(api_filter_id: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApiFilterResult:
    """
    Supported in provider *v3.14+* and VCD 10.4.3+.

    Provides a data source to read API Filters in VMware Cloud Director. An API Filter allows to extend VCD API with customised URLs
    that can be redirected to an [`ExternalEndpoint`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/external_endpoint).

    > Only `System Administrator` can use this data source.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    api_filter1 = vcd.get_api_filter(api_filter_id="urn:vcloud:apiFilter:4252ab09-eed8-4bc6-86d7-6019090273f5")
    ```


    :param str api_filter_id: ID of the API Filter. This is the only way of unequivocally identify an API Filter. A list of
           available API Filters can be obtained by using the `list@` option of the import mechanism of the [resource](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/api_filter#importing)
    """
    __args__ = dict()
    __args__['apiFilterId'] = api_filter_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getApiFilter:getApiFilter', __args__, opts=opts, typ=GetApiFilterResult).value

    return AwaitableGetApiFilterResult(
        api_filter_id=pulumi.get(__ret__, 'api_filter_id'),
        external_endpoint_id=pulumi.get(__ret__, 'external_endpoint_id'),
        id=pulumi.get(__ret__, 'id'),
        url_matcher_pattern=pulumi.get(__ret__, 'url_matcher_pattern'),
        url_matcher_scope=pulumi.get(__ret__, 'url_matcher_scope'))
def get_api_filter_output(api_filter_id: Optional[pulumi.Input[str]] = None,
                          opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetApiFilterResult]:
    """
    Supported in provider *v3.14+* and VCD 10.4.3+.

    Provides a data source to read API Filters in VMware Cloud Director. An API Filter allows to extend VCD API with customised URLs
    that can be redirected to an [`ExternalEndpoint`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/external_endpoint).

    > Only `System Administrator` can use this data source.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    api_filter1 = vcd.get_api_filter(api_filter_id="urn:vcloud:apiFilter:4252ab09-eed8-4bc6-86d7-6019090273f5")
    ```


    :param str api_filter_id: ID of the API Filter. This is the only way of unequivocally identify an API Filter. A list of
           available API Filters can be obtained by using the `list@` option of the import mechanism of the [resource](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/api_filter#importing)
    """
    __args__ = dict()
    __args__['apiFilterId'] = api_filter_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getApiFilter:getApiFilter', __args__, opts=opts, typ=GetApiFilterResult)
    return __ret__.apply(lambda __response__: GetApiFilterResult(
        api_filter_id=pulumi.get(__response__, 'api_filter_id'),
        external_endpoint_id=pulumi.get(__response__, 'external_endpoint_id'),
        id=pulumi.get(__response__, 'id'),
        url_matcher_pattern=pulumi.get(__response__, 'url_matcher_pattern'),
        url_matcher_scope=pulumi.get(__response__, 'url_matcher_scope')))
