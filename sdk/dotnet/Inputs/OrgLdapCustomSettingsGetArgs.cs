// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class OrgLdapCustomSettingsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// authentication method: one of SIMPLE, MD5DIGEST, NTLM
        /// </summary>
        [Input("authenticationMethod", required: true)]
        public Input<string> AuthenticationMethod { get; set; } = null!;

        /// <summary>
        /// LDAP search base
        /// </summary>
        [Input("baseDistinguishedName")]
        public Input<string>? BaseDistinguishedName { get; set; }

        /// <summary>
        /// type of connector: one of OPEN_LDAP, ACTIVE_DIRECTORY
        /// </summary>
        [Input("connectorType", required: true)]
        public Input<string> ConnectorType { get; set; } = null!;

        /// <summary>
        /// Group settings when `ldap_mode` is CUSTOM
        /// </summary>
        [Input("groupAttributes", required: true)]
        public Input<Inputs.OrgLdapCustomSettingsGroupAttributesGetArgs> GroupAttributes { get; set; } = null!;

        /// <summary>
        /// True if the LDAP service requires an SSL connection
        /// </summary>
        [Input("isSsl")]
        public Input<bool>? IsSsl { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for the user identified by UserName. This value is never returned by GET. It is inspected on create and modify. On modify, the absence of this element indicates that the password should not be changed
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Port number for LDAP service
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// host name or IP of the LDAP server
        /// </summary>
        [Input("server", required: true)]
        public Input<string> Server { get; set; } = null!;

        /// <summary>
        /// User settings when `ldap_mode` is CUSTOM
        /// </summary>
        [Input("userAttributes", required: true)]
        public Input<Inputs.OrgLdapCustomSettingsUserAttributesGetArgs> UserAttributes { get; set; } = null!;

        /// <summary>
        /// Username to use when logging in to LDAP, specified using LDAP attribute=value pairs (for example: cn="ldap-admin", c="example", dc="com")
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public OrgLdapCustomSettingsGetArgs()
        {
        }
        public static new OrgLdapCustomSettingsGetArgs Empty => new OrgLdapCustomSettingsGetArgs();
    }
}
