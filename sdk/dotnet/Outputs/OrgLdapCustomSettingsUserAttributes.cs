// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Outputs
{

    [OutputType]
    public sealed class OrgLdapCustomSettingsUserAttributes
    {
        /// <summary>
        /// LDAP attribute to use for the user's full name. For example, displayName
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// LDAP attribute to use for the user's email address. For example, mail
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// LDAP attribute to use for the user's given name. For example, givenName
        /// </summary>
        public readonly string GivenName;
        /// <summary>
        /// LDAP attribute that returns the identifiers of all the groups of which the user is a member
        /// </summary>
        public readonly string? GroupBackLinkIdentifier;
        /// <summary>
        /// LDAP attribute that identifies a user as a member of a group. For example, dn
        /// </summary>
        public readonly string GroupMembershipIdentifier;
        /// <summary>
        /// LDAP objectClass of which imported users are members. For example, user or person
        /// </summary>
        public readonly string ObjectClass;
        /// <summary>
        /// LDAP attribute to use for the user's surname. For example, sn
        /// </summary>
        public readonly string Surname;
        /// <summary>
        /// LDAP attribute to use for the user's telephone number. For example, telephoneNumber
        /// </summary>
        public readonly string Telephone;
        /// <summary>
        /// LDAP attribute to use as the unique identifier for a user. For example, objectGuid
        /// </summary>
        public readonly string UniqueIdentifier;
        /// <summary>
        /// LDAP attribute to use when looking up a user name to import. For example, userPrincipalName or samAccountName
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private OrgLdapCustomSettingsUserAttributes(
            string displayName,

            string email,

            string givenName,

            string? groupBackLinkIdentifier,

            string groupMembershipIdentifier,

            string objectClass,

            string surname,

            string telephone,

            string uniqueIdentifier,

            string username)
        {
            DisplayName = displayName;
            Email = email;
            GivenName = givenName;
            GroupBackLinkIdentifier = groupBackLinkIdentifier;
            GroupMembershipIdentifier = groupMembershipIdentifier;
            ObjectClass = objectClass;
            Surname = surname;
            Telephone = telephone;
            UniqueIdentifier = uniqueIdentifier;
            Username = username;
        }
    }
}
