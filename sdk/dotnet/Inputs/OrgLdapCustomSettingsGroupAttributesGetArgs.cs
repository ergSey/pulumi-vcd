// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class OrgLdapCustomSettingsGroupAttributesGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// LDAP group attribute used to identify a group member
        /// </summary>
        [Input("groupBackLinkIdentifier")]
        public Input<string>? GroupBackLinkIdentifier { get; set; }

        /// <summary>
        /// LDAP attribute that identifies a group as a member of another group. For example, dn
        /// </summary>
        [Input("groupMembershipIdentifier", required: true)]
        public Input<string> GroupMembershipIdentifier { get; set; } = null!;

        /// <summary>
        /// LDAP attribute to use when getting the members of a group. For example, member
        /// </summary>
        [Input("membership", required: true)]
        public Input<string> Membership { get; set; } = null!;

        /// <summary>
        /// LDAP attribute to use for the group name. For example, cn
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// LDAP objectClass of which imported groups are members. For example, group
        /// </summary>
        [Input("objectClass", required: true)]
        public Input<string> ObjectClass { get; set; } = null!;

        /// <summary>
        /// LDAP attribute to use as the unique identifier for a group. For example, objectGuid
        /// </summary>
        [Input("uniqueIdentifier", required: true)]
        public Input<string> UniqueIdentifier { get; set; } = null!;

        public OrgLdapCustomSettingsGroupAttributesGetArgs()
        {
        }
        public static new OrgLdapCustomSettingsGroupAttributesGetArgs Empty => new OrgLdapCustomSettingsGroupAttributesGetArgs();
    }
}
