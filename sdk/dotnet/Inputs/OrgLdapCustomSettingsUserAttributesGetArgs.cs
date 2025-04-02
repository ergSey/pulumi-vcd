// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class OrgLdapCustomSettingsUserAttributesGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// LDAP attribute to use for the user's full name. For example, displayName
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// LDAP attribute to use for the user's email address. For example, mail
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        /// <summary>
        /// LDAP attribute to use for the user's given name. For example, givenName
        /// </summary>
        [Input("givenName", required: true)]
        public Input<string> GivenName { get; set; } = null!;

        /// <summary>
        /// LDAP attribute that returns the identifiers of all the groups of which the user is a member
        /// </summary>
        [Input("groupBackLinkIdentifier")]
        public Input<string>? GroupBackLinkIdentifier { get; set; }

        /// <summary>
        /// LDAP attribute that identifies a user as a member of a group. For example, dn
        /// </summary>
        [Input("groupMembershipIdentifier", required: true)]
        public Input<string> GroupMembershipIdentifier { get; set; } = null!;

        /// <summary>
        /// LDAP objectClass of which imported users are members. For example, user or person
        /// </summary>
        [Input("objectClass", required: true)]
        public Input<string> ObjectClass { get; set; } = null!;

        /// <summary>
        /// LDAP attribute to use for the user's surname. For example, sn
        /// </summary>
        [Input("surname", required: true)]
        public Input<string> Surname { get; set; } = null!;

        /// <summary>
        /// LDAP attribute to use for the user's telephone number. For example, telephoneNumber
        /// </summary>
        [Input("telephone", required: true)]
        public Input<string> Telephone { get; set; } = null!;

        /// <summary>
        /// LDAP attribute to use as the unique identifier for a user. For example, objectGuid
        /// </summary>
        [Input("uniqueIdentifier", required: true)]
        public Input<string> UniqueIdentifier { get; set; } = null!;

        /// <summary>
        /// LDAP attribute to use when looking up a user name to import. For example, userPrincipalName or samAccountName
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public OrgLdapCustomSettingsUserAttributesGetArgs()
        {
        }
        public static new OrgLdapCustomSettingsUserAttributesGetArgs Empty => new OrgLdapCustomSettingsUserAttributesGetArgs();
    }
}
