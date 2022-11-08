// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.CloudTasks.V2Beta2
{
    /// <summary>
    /// The HTTP method to use for the request. The default is POST. The app's request handler for the task's target URL must be able to handle HTTP requests with this http_method, otherwise the task attempt fails with error code 405 (Method Not Allowed). See [Writing a push task request handler](https://cloud.google.com/appengine/docs/java/taskqueue/push/creating-handlers#writing_a_push_task_request_handler) and the App Engine documentation for your runtime on [How Requests are Handled](https://cloud.google.com/appengine/docs/standard/python3/how-requests-are-handled).
    /// </summary>
    [EnumType]
    public readonly struct AppEngineHttpRequestHttpMethod : IEquatable<AppEngineHttpRequestHttpMethod>
    {
        private readonly string _value;

        private AppEngineHttpRequestHttpMethod(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// HTTP method unspecified
        /// </summary>
        public static AppEngineHttpRequestHttpMethod HttpMethodUnspecified { get; } = new AppEngineHttpRequestHttpMethod("HTTP_METHOD_UNSPECIFIED");
        /// <summary>
        /// HTTP POST
        /// </summary>
        public static AppEngineHttpRequestHttpMethod Post { get; } = new AppEngineHttpRequestHttpMethod("POST");
        /// <summary>
        /// HTTP GET
        /// </summary>
        public static AppEngineHttpRequestHttpMethod Get { get; } = new AppEngineHttpRequestHttpMethod("GET");
        /// <summary>
        /// HTTP HEAD
        /// </summary>
        public static AppEngineHttpRequestHttpMethod Head { get; } = new AppEngineHttpRequestHttpMethod("HEAD");
        /// <summary>
        /// HTTP PUT
        /// </summary>
        public static AppEngineHttpRequestHttpMethod Put { get; } = new AppEngineHttpRequestHttpMethod("PUT");
        /// <summary>
        /// HTTP DELETE
        /// </summary>
        public static AppEngineHttpRequestHttpMethod Delete { get; } = new AppEngineHttpRequestHttpMethod("DELETE");
        /// <summary>
        /// HTTP PATCH
        /// </summary>
        public static AppEngineHttpRequestHttpMethod Patch { get; } = new AppEngineHttpRequestHttpMethod("PATCH");
        /// <summary>
        /// HTTP OPTIONS
        /// </summary>
        public static AppEngineHttpRequestHttpMethod Options { get; } = new AppEngineHttpRequestHttpMethod("OPTIONS");

        public static bool operator ==(AppEngineHttpRequestHttpMethod left, AppEngineHttpRequestHttpMethod right) => left.Equals(right);
        public static bool operator !=(AppEngineHttpRequestHttpMethod left, AppEngineHttpRequestHttpMethod right) => !left.Equals(right);

        public static explicit operator string(AppEngineHttpRequestHttpMethod value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AppEngineHttpRequestHttpMethod other && Equals(other);
        public bool Equals(AppEngineHttpRequestHttpMethod other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The HTTP method to use for the request. The default is POST.
    /// </summary>
    [EnumType]
    public readonly struct HttpRequestHttpMethod : IEquatable<HttpRequestHttpMethod>
    {
        private readonly string _value;

        private HttpRequestHttpMethod(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// HTTP method unspecified
        /// </summary>
        public static HttpRequestHttpMethod HttpMethodUnspecified { get; } = new HttpRequestHttpMethod("HTTP_METHOD_UNSPECIFIED");
        /// <summary>
        /// HTTP POST
        /// </summary>
        public static HttpRequestHttpMethod Post { get; } = new HttpRequestHttpMethod("POST");
        /// <summary>
        /// HTTP GET
        /// </summary>
        public static HttpRequestHttpMethod Get { get; } = new HttpRequestHttpMethod("GET");
        /// <summary>
        /// HTTP HEAD
        /// </summary>
        public static HttpRequestHttpMethod Head { get; } = new HttpRequestHttpMethod("HEAD");
        /// <summary>
        /// HTTP PUT
        /// </summary>
        public static HttpRequestHttpMethod Put { get; } = new HttpRequestHttpMethod("PUT");
        /// <summary>
        /// HTTP DELETE
        /// </summary>
        public static HttpRequestHttpMethod Delete { get; } = new HttpRequestHttpMethod("DELETE");
        /// <summary>
        /// HTTP PATCH
        /// </summary>
        public static HttpRequestHttpMethod Patch { get; } = new HttpRequestHttpMethod("PATCH");
        /// <summary>
        /// HTTP OPTIONS
        /// </summary>
        public static HttpRequestHttpMethod Options { get; } = new HttpRequestHttpMethod("OPTIONS");

        public static bool operator ==(HttpRequestHttpMethod left, HttpRequestHttpMethod right) => left.Equals(right);
        public static bool operator !=(HttpRequestHttpMethod left, HttpRequestHttpMethod right) => !left.Equals(right);

        public static explicit operator string(HttpRequestHttpMethod value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is HttpRequestHttpMethod other && Equals(other);
        public bool Equals(HttpRequestHttpMethod other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The HTTP method to use for the request. When specified, it will override HttpRequest for the task. Note that if the value is set to HttpMethod the HttpRequest of the task will be ignored at execution time.
    /// </summary>
    [EnumType]
    public readonly struct HttpTargetHttpMethod : IEquatable<HttpTargetHttpMethod>
    {
        private readonly string _value;

        private HttpTargetHttpMethod(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// HTTP method unspecified
        /// </summary>
        public static HttpTargetHttpMethod HttpMethodUnspecified { get; } = new HttpTargetHttpMethod("HTTP_METHOD_UNSPECIFIED");
        /// <summary>
        /// HTTP POST
        /// </summary>
        public static HttpTargetHttpMethod Post { get; } = new HttpTargetHttpMethod("POST");
        /// <summary>
        /// HTTP GET
        /// </summary>
        public static HttpTargetHttpMethod Get { get; } = new HttpTargetHttpMethod("GET");
        /// <summary>
        /// HTTP HEAD
        /// </summary>
        public static HttpTargetHttpMethod Head { get; } = new HttpTargetHttpMethod("HEAD");
        /// <summary>
        /// HTTP PUT
        /// </summary>
        public static HttpTargetHttpMethod Put { get; } = new HttpTargetHttpMethod("PUT");
        /// <summary>
        /// HTTP DELETE
        /// </summary>
        public static HttpTargetHttpMethod Delete { get; } = new HttpTargetHttpMethod("DELETE");
        /// <summary>
        /// HTTP PATCH
        /// </summary>
        public static HttpTargetHttpMethod Patch { get; } = new HttpTargetHttpMethod("PATCH");
        /// <summary>
        /// HTTP OPTIONS
        /// </summary>
        public static HttpTargetHttpMethod Options { get; } = new HttpTargetHttpMethod("OPTIONS");

        public static bool operator ==(HttpTargetHttpMethod left, HttpTargetHttpMethod right) => left.Equals(right);
        public static bool operator !=(HttpTargetHttpMethod left, HttpTargetHttpMethod right) => !left.Equals(right);

        public static explicit operator string(HttpTargetHttpMethod value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is HttpTargetHttpMethod other && Equals(other);
        public bool Equals(HttpTargetHttpMethod other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The response_view specifies which subset of the Task will be returned. By default response_view is BASIC; not all information is retrieved by default because some data, such as payloads, might be desirable to return only when needed because of its large size or because of the sensitivity of data that it contains. Authorization for FULL requires `cloudtasks.tasks.fullView` [Google IAM](https://cloud.google.com/iam/) permission on the Task resource.
    /// </summary>
    [EnumType]
    public readonly struct TaskResponseView : IEquatable<TaskResponseView>
    {
        private readonly string _value;

        private TaskResponseView(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Unspecified. Defaults to BASIC.
        /// </summary>
        public static TaskResponseView ViewUnspecified { get; } = new TaskResponseView("VIEW_UNSPECIFIED");
        /// <summary>
        /// The basic view omits fields which can be large or can contain sensitive data. This view does not include the (payload in AppEngineHttpRequest and payload in PullMessage). These payloads are desirable to return only when needed, because they can be large and because of the sensitivity of the data that you choose to store in it.
        /// </summary>
        public static TaskResponseView Basic { get; } = new TaskResponseView("BASIC");
        /// <summary>
        /// All information is returned. Authorization for FULL requires `cloudtasks.tasks.fullView` [Google IAM](https://cloud.google.com/iam/) permission on the Queue resource.
        /// </summary>
        public static TaskResponseView Full { get; } = new TaskResponseView("FULL");

        public static bool operator ==(TaskResponseView left, TaskResponseView right) => left.Equals(right);
        public static bool operator !=(TaskResponseView left, TaskResponseView right) => !left.Equals(right);

        public static explicit operator string(TaskResponseView value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskResponseView other && Equals(other);
        public bool Equals(TaskResponseView other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Scheme override. When specified, the Uri scheme is replaced by the provided value.
    /// </summary>
    [EnumType]
    public readonly struct UriOverrideScheme : IEquatable<UriOverrideScheme>
    {
        private readonly string _value;

        private UriOverrideScheme(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Scheme unspecified. Defaults to HTTPS.
        /// </summary>
        public static UriOverrideScheme SchemeUnspecified { get; } = new UriOverrideScheme("SCHEME_UNSPECIFIED");
        /// <summary>
        /// Convert the scheme to HTTP, e.g., https://www.google.ca will change to http://www.google.ca.
        /// </summary>
        public static UriOverrideScheme Http { get; } = new UriOverrideScheme("HTTP");
        /// <summary>
        /// Convert the scheme to HTTPS, e.g., http://www.google.ca will change to https://www.google.ca.
        /// </summary>
        public static UriOverrideScheme Https { get; } = new UriOverrideScheme("HTTPS");

        public static bool operator ==(UriOverrideScheme left, UriOverrideScheme right) => left.Equals(right);
        public static bool operator !=(UriOverrideScheme left, UriOverrideScheme right) => !left.Equals(right);

        public static explicit operator string(UriOverrideScheme value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is UriOverrideScheme other && Equals(other);
        public bool Equals(UriOverrideScheme other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Uri Override Enforce Mode Determines the Target UriOverride mode.
    /// </summary>
    [EnumType]
    public readonly struct UriOverrideUriOverrideEnforceMode : IEquatable<UriOverrideUriOverrideEnforceMode>
    {
        private readonly string _value;

        private UriOverrideUriOverrideEnforceMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// OverrideMode Unspecified. Defaults to ALWAYS.
        /// </summary>
        public static UriOverrideUriOverrideEnforceMode UriOverrideEnforceModeUnspecified { get; } = new UriOverrideUriOverrideEnforceMode("URI_OVERRIDE_ENFORCE_MODE_UNSPECIFIED");
        /// <summary>
        /// In the IF_NOT_EXISTS mode, queue-level configuration is only applied where task-level configuration does not exist.
        /// </summary>
        public static UriOverrideUriOverrideEnforceMode IfNotExists { get; } = new UriOverrideUriOverrideEnforceMode("IF_NOT_EXISTS");
        /// <summary>
        /// In the ALWAYS mode, queue-level configuration overrides all task-level configuration
        /// </summary>
        public static UriOverrideUriOverrideEnforceMode Always { get; } = new UriOverrideUriOverrideEnforceMode("ALWAYS");

        public static bool operator ==(UriOverrideUriOverrideEnforceMode left, UriOverrideUriOverrideEnforceMode right) => left.Equals(right);
        public static bool operator !=(UriOverrideUriOverrideEnforceMode left, UriOverrideUriOverrideEnforceMode right) => !left.Equals(right);

        public static explicit operator string(UriOverrideUriOverrideEnforceMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is UriOverrideUriOverrideEnforceMode other && Equals(other);
        public bool Equals(UriOverrideUriOverrideEnforceMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
