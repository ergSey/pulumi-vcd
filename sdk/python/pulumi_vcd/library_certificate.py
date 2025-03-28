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

__all__ = ['LibraryCertificateArgs', 'LibraryCertificate']

@pulumi.input_type
class LibraryCertificateArgs:
    def __init__(__self__, *,
                 alias: pulumi.Input[str],
                 certificate: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 private_key_passphrase: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LibraryCertificate resource.
        :param pulumi.Input[str] alias: Alias of certificate
        :param pulumi.Input[str] certificate: Certificate content
        :param pulumi.Input[str] description: Certificate description
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[str] private_key: Certificate private key
        :param pulumi.Input[str] private_key_passphrase: Certificate private pass phrase
        """
        pulumi.set(__self__, "alias", alias)
        pulumi.set(__self__, "certificate", certificate)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)
        if private_key_passphrase is not None:
            pulumi.set(__self__, "private_key_passphrase", private_key_passphrase)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Input[str]:
        """
        Alias of certificate
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: pulumi.Input[str]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Input[str]:
        """
        Certificate content
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: pulumi.Input[str]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def org(self) -> Optional[pulumi.Input[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        different organizations
        """
        return pulumi.get(self, "org")

    @org.setter
    def org(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate private key
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter(name="privateKeyPassphrase")
    def private_key_passphrase(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate private pass phrase
        """
        return pulumi.get(self, "private_key_passphrase")

    @private_key_passphrase.setter
    def private_key_passphrase(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key_passphrase", value)


@pulumi.input_type
class _LibraryCertificateState:
    def __init__(__self__, *,
                 alias: Optional[pulumi.Input[str]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 private_key_passphrase: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LibraryCertificate resources.
        :param pulumi.Input[str] alias: Alias of certificate
        :param pulumi.Input[str] certificate: Certificate content
        :param pulumi.Input[str] description: Certificate description
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[str] private_key: Certificate private key
        :param pulumi.Input[str] private_key_passphrase: Certificate private pass phrase
        """
        if alias is not None:
            pulumi.set(__self__, "alias", alias)
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)
        if private_key_passphrase is not None:
            pulumi.set(__self__, "private_key_passphrase", private_key_passphrase)

    @property
    @pulumi.getter
    def alias(self) -> Optional[pulumi.Input[str]]:
        """
        Alias of certificate
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate content
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def org(self) -> Optional[pulumi.Input[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        different organizations
        """
        return pulumi.get(self, "org")

    @org.setter
    def org(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate private key
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter(name="privateKeyPassphrase")
    def private_key_passphrase(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate private pass phrase
        """
        return pulumi.get(self, "private_key_passphrase")

    @private_key_passphrase.setter
    def private_key_passphrase(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key_passphrase", value)


class LibraryCertificate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 private_key_passphrase: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a LibraryCertificate resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: Alias of certificate
        :param pulumi.Input[str] certificate: Certificate content
        :param pulumi.Input[str] description: Certificate description
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[str] private_key: Certificate private key
        :param pulumi.Input[str] private_key_passphrase: Certificate private pass phrase
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LibraryCertificateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a LibraryCertificate resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param LibraryCertificateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LibraryCertificateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 private_key_passphrase: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LibraryCertificateArgs.__new__(LibraryCertificateArgs)

            if alias is None and not opts.urn:
                raise TypeError("Missing required property 'alias'")
            __props__.__dict__["alias"] = alias
            if certificate is None and not opts.urn:
                raise TypeError("Missing required property 'certificate'")
            __props__.__dict__["certificate"] = certificate
            __props__.__dict__["description"] = description
            __props__.__dict__["org"] = org
            __props__.__dict__["private_key"] = None if private_key is None else pulumi.Output.secret(private_key)
            __props__.__dict__["private_key_passphrase"] = None if private_key_passphrase is None else pulumi.Output.secret(private_key_passphrase)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["privateKey", "privateKeyPassphrase"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(LibraryCertificate, __self__).__init__(
            'vcd:index/libraryCertificate:LibraryCertificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alias: Optional[pulumi.Input[str]] = None,
            certificate: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            org: Optional[pulumi.Input[str]] = None,
            private_key: Optional[pulumi.Input[str]] = None,
            private_key_passphrase: Optional[pulumi.Input[str]] = None) -> 'LibraryCertificate':
        """
        Get an existing LibraryCertificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: Alias of certificate
        :param pulumi.Input[str] certificate: Certificate content
        :param pulumi.Input[str] description: Certificate description
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[str] private_key: Certificate private key
        :param pulumi.Input[str] private_key_passphrase: Certificate private pass phrase
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LibraryCertificateState.__new__(_LibraryCertificateState)

        __props__.__dict__["alias"] = alias
        __props__.__dict__["certificate"] = certificate
        __props__.__dict__["description"] = description
        __props__.__dict__["org"] = org
        __props__.__dict__["private_key"] = private_key
        __props__.__dict__["private_key_passphrase"] = private_key_passphrase
        return LibraryCertificate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Output[str]:
        """
        Alias of certificate
        """
        return pulumi.get(self, "alias")

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Output[str]:
        """
        Certificate content
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Certificate description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def org(self) -> pulumi.Output[Optional[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        different organizations
        """
        return pulumi.get(self, "org")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Output[Optional[str]]:
        """
        Certificate private key
        """
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter(name="privateKeyPassphrase")
    def private_key_passphrase(self) -> pulumi.Output[Optional[str]]:
        """
        Certificate private pass phrase
        """
        return pulumi.get(self, "private_key_passphrase")

