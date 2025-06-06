// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Vcd
{
    public static class Config
    {
        [global::System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("vcd");

        private static readonly __Value<bool?> _allowApiTokenFile = new __Value<bool?>(() => __config.GetBoolean("allowApiTokenFile"));
        /// <summary>
        /// Set this to true if you understand the security risks of using API token files and would like to suppress the warnings
        /// </summary>
        public static bool? AllowApiTokenFile
        {
            get => _allowApiTokenFile.Get();
            set => _allowApiTokenFile.Set(value);
        }

        private static readonly __Value<bool?> _allowServiceAccountTokenFile = new __Value<bool?>(() => __config.GetBoolean("allowServiceAccountTokenFile"));
        /// <summary>
        /// Set this to true if you understand the security risks of using Service Account token files and would like to suppress
        /// the warnings
        /// </summary>
        public static bool? AllowServiceAccountTokenFile
        {
            get => _allowServiceAccountTokenFile.Get();
            set => _allowServiceAccountTokenFile.Set(value);
        }

        private static readonly __Value<bool?> _allowUnverifiedSsl = new __Value<bool?>(() => __config.GetBoolean("allowUnverifiedSsl"));
        /// <summary>
        /// If set, VCDClient will permit unverifiable SSL certificates.
        /// </summary>
        public static bool? AllowUnverifiedSsl
        {
            get => _allowUnverifiedSsl.Get();
            set => _allowUnverifiedSsl.Set(value);
        }

        private static readonly __Value<string?> _apiToken = new __Value<string?>(() => __config.Get("apiToken"));
        /// <summary>
        /// The API token used instead of username/password for VCD API operations. (Requires VCD 10.3.1+)
        /// </summary>
        public static string? ApiToken
        {
            get => _apiToken.Get();
            set => _apiToken.Set(value);
        }

        private static readonly __Value<string?> _apiTokenFile = new __Value<string?>(() => __config.Get("apiTokenFile"));
        /// <summary>
        /// The API token file instead of username/password for VCD API operations. (Requires VCD 10.3.1+)
        /// </summary>
        public static string? ApiTokenFile
        {
            get => _apiTokenFile.Get();
            set => _apiTokenFile.Set(value);
        }

        private static readonly __Value<string?> _authType = new __Value<string?>(() => __config.Get("authType"));
        /// <summary>
        /// 'integrated', 'saml_adfs', 'token', 'api_token', 'api_token_file' and 'service_account_token_file' are supported.
        /// 'integrated' is default.
        /// </summary>
        public static string? AuthType
        {
            get => _authType.Get();
            set => _authType.Set(value);
        }

        private static readonly __Value<ImmutableArray<Pulumi.Vcd.Config.Types.IgnoreMetadataChanges>> _ignoreMetadataChanges = new __Value<ImmutableArray<Pulumi.Vcd.Config.Types.IgnoreMetadataChanges>>(() => __config.GetObject<ImmutableArray<Pulumi.Vcd.Config.Types.IgnoreMetadataChanges>>("ignoreMetadataChanges"));
        /// <summary>
        /// Defines a set of `metadata_entry` that need to be ignored by this provider. All filters on this attribute are computed
        /// with a logical AND
        /// </summary>
        public static ImmutableArray<Pulumi.Vcd.Config.Types.IgnoreMetadataChanges> IgnoreMetadataChanges
        {
            get => _ignoreMetadataChanges.Get();
            set => _ignoreMetadataChanges.Set(value);
        }

        private static readonly __Value<string?> _importSeparator = new __Value<string?>(() => __config.Get("importSeparator"));
        public static string? ImportSeparator
        {
            get => _importSeparator.Get();
            set => _importSeparator.Set(value);
        }

        private static readonly __Value<bool?> _logging = new __Value<bool?>(() => __config.GetBoolean("logging"));
        /// <summary>
        /// If set, it will enable logging of API requests and responses
        /// </summary>
        public static bool? Logging
        {
            get => _logging.Get();
            set => _logging.Set(value);
        }

        private static readonly __Value<string?> _loggingFile = new __Value<string?>(() => __config.Get("loggingFile"));
        /// <summary>
        /// Defines the full name of the logging file for API calls (requires 'logging')
        /// </summary>
        public static string? LoggingFile
        {
            get => _loggingFile.Get();
            set => _loggingFile.Set(value);
        }

        private static readonly __Value<int?> _maxRetryTimeout = new __Value<int?>(() => __config.GetInt32("maxRetryTimeout"));
        /// <summary>
        /// Max num seconds to wait for successful response when operating on resources within vCloud (defaults to 60)
        /// </summary>
        public static int? MaxRetryTimeout
        {
            get => _maxRetryTimeout.Get();
            set => _maxRetryTimeout.Set(value);
        }

        private static readonly __Value<string?> _org = new __Value<string?>(() => __config.Get("org"));
        /// <summary>
        /// The VCD Org for API operations
        /// </summary>
        public static string? Org
        {
            get => _org.Get();
            set => _org.Set(value);
        }

        private static readonly __Value<string?> _password = new __Value<string?>(() => __config.Get("password"));
        /// <summary>
        /// The user password for VCD API operations.
        /// </summary>
        public static string? Password
        {
            get => _password.Get();
            set => _password.Set(value);
        }

        private static readonly __Value<string?> _samlAdfsCookie = new __Value<string?>(() => __config.Get("samlAdfsCookie"));
        /// <summary>
        /// Allows to specify custom cookie for ADFS server lookup. '{{.Org}}' is replaced by real Org - e.g. 'sso-preferred=yes;
        /// sso_redirect_org={{.Org}}'
        /// </summary>
        public static string? SamlAdfsCookie
        {
            get => _samlAdfsCookie.Get();
            set => _samlAdfsCookie.Set(value);
        }

        private static readonly __Value<string?> _samlAdfsRptId = new __Value<string?>(() => __config.Get("samlAdfsRptId"));
        /// <summary>
        /// Allows to specify custom Relaying Party Trust Identifier for auth_type=saml_adfs
        /// </summary>
        public static string? SamlAdfsRptId
        {
            get => _samlAdfsRptId.Get();
            set => _samlAdfsRptId.Set(value);
        }

        private static readonly __Value<string?> _serviceAccountTokenFile = new __Value<string?>(() => __config.Get("serviceAccountTokenFile"));
        /// <summary>
        /// The Service Account API token file instead of username/password for VCD API operations. (Requires VCD 10.4.0+)
        /// </summary>
        public static string? ServiceAccountTokenFile
        {
            get => _serviceAccountTokenFile.Get();
            set => _serviceAccountTokenFile.Set(value);
        }

        private static readonly __Value<string?> _sysorg = new __Value<string?>(() => __config.Get("sysorg"));
        /// <summary>
        /// The VCD Org for user authentication
        /// </summary>
        public static string? Sysorg
        {
            get => _sysorg.Get();
            set => _sysorg.Set(value);
        }

        private static readonly __Value<string?> _token = new __Value<string?>(() => __config.Get("token"));
        /// <summary>
        /// The token used instead of username/password for VCD API operations.
        /// </summary>
        public static string? Token
        {
            get => _token.Get();
            set => _token.Set(value);
        }

        private static readonly __Value<string?> _url = new __Value<string?>(() => __config.Get("url"));
        /// <summary>
        /// The VCD url for VCD API operations.
        /// </summary>
        public static string? Url
        {
            get => _url.Get();
            set => _url.Set(value);
        }

        private static readonly __Value<string?> _user = new __Value<string?>(() => __config.Get("user"));
        /// <summary>
        /// The user name for VCD API operations.
        /// </summary>
        public static string? User
        {
            get => _user.Get();
            set => _user.Set(value);
        }

        private static readonly __Value<string?> _vdc = new __Value<string?>(() => __config.Get("vdc"));
        /// <summary>
        /// The VDC for API operations
        /// </summary>
        public static string? Vdc
        {
            get => _vdc.Get();
            set => _vdc.Set(value);
        }

        public static class Types
        {

             public class IgnoreMetadataChanges
             {
                public string? ConflictAction { get; set; } = null!;
            /// <summary>
            /// Regular expression of the metadata entry keys to ignore. Either `key_regex` or `value_regex` is required
            /// </summary>
                public string? KeyRegex { get; set; } = null!;
            /// <summary>
            /// Ignores metadata from the specific entity in VCD named like this argument
            /// </summary>
                public string? ResourceName { get; set; } = null!;
            /// <summary>
            /// Ignores metadata from the specific resource type
            /// </summary>
                public string? ResourceType { get; set; } = null!;
            /// <summary>
            /// Regular expression of the metadata entry values to ignore. Either `key_regex` or `value_regex` is required
            /// </summary>
                public string? ValueRegex { get; set; } = null!;
            }
        }
    }
}
