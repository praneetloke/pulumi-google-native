# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AppEngineHttpRequestHttpMethod',
    'HttpRequestHttpMethod',
    'HttpTargetHttpMethod',
    'QueueType',
    'TaskResponseView',
    'UriOverrideScheme',
    'UriOverrideUriOverrideEnforceMode',
]


class AppEngineHttpRequestHttpMethod(str, Enum):
    """
    The HTTP method to use for the request. The default is POST. The app's request handler for the task's target URL must be able to handle HTTP requests with this http_method, otherwise the task attempt fails with error code 405 (Method Not Allowed). See [Writing a push task request handler](https://cloud.google.com/appengine/docs/java/taskqueue/push/creating-handlers#writing_a_push_task_request_handler) and the App Engine documentation for your runtime on [How Requests are Handled](https://cloud.google.com/appengine/docs/standard/python3/how-requests-are-handled).
    """
    HTTP_METHOD_UNSPECIFIED = "HTTP_METHOD_UNSPECIFIED"
    """
    HTTP method unspecified
    """
    POST = "POST"
    """
    HTTP POST
    """
    GET = "GET"
    """
    HTTP GET
    """
    HEAD = "HEAD"
    """
    HTTP HEAD
    """
    PUT = "PUT"
    """
    HTTP PUT
    """
    DELETE = "DELETE"
    """
    HTTP DELETE
    """
    PATCH = "PATCH"
    """
    HTTP PATCH
    """
    OPTIONS = "OPTIONS"
    """
    HTTP OPTIONS
    """


class HttpRequestHttpMethod(str, Enum):
    """
    The HTTP method to use for the request. The default is POST.
    """
    HTTP_METHOD_UNSPECIFIED = "HTTP_METHOD_UNSPECIFIED"
    """
    HTTP method unspecified
    """
    POST = "POST"
    """
    HTTP POST
    """
    GET = "GET"
    """
    HTTP GET
    """
    HEAD = "HEAD"
    """
    HTTP HEAD
    """
    PUT = "PUT"
    """
    HTTP PUT
    """
    DELETE = "DELETE"
    """
    HTTP DELETE
    """
    PATCH = "PATCH"
    """
    HTTP PATCH
    """
    OPTIONS = "OPTIONS"
    """
    HTTP OPTIONS
    """


class HttpTargetHttpMethod(str, Enum):
    """
    The HTTP method to use for the request. When specified, it overrides HttpRequest for the task. Note that if the value is set to HttpMethod the HttpRequest of the task will be ignored at execution time.
    """
    HTTP_METHOD_UNSPECIFIED = "HTTP_METHOD_UNSPECIFIED"
    """
    HTTP method unspecified
    """
    POST = "POST"
    """
    HTTP POST
    """
    GET = "GET"
    """
    HTTP GET
    """
    HEAD = "HEAD"
    """
    HTTP HEAD
    """
    PUT = "PUT"
    """
    HTTP PUT
    """
    DELETE = "DELETE"
    """
    HTTP DELETE
    """
    PATCH = "PATCH"
    """
    HTTP PATCH
    """
    OPTIONS = "OPTIONS"
    """
    HTTP OPTIONS
    """


class QueueType(str, Enum):
    """
    Immutable. The type of a queue (push or pull). `Queue.type` is an immutable property of the queue that is set at the queue creation time. When left unspecified, the default value of `PUSH` is selected.
    """
    TYPE_UNSPECIFIED = "TYPE_UNSPECIFIED"
    """
    Default value.
    """
    PULL = "PULL"
    """
    A pull queue.
    """
    PUSH = "PUSH"
    """
    A push queue.
    """


class TaskResponseView(str, Enum):
    """
    The response_view specifies which subset of the Task will be returned. By default response_view is BASIC; not all information is retrieved by default because some data, such as payloads, might be desirable to return only when needed because of its large size or because of the sensitivity of data that it contains. Authorization for FULL requires `cloudtasks.tasks.fullView` [Google IAM](https://cloud.google.com/iam/) permission on the Task resource.
    """
    VIEW_UNSPECIFIED = "VIEW_UNSPECIFIED"
    """
    Unspecified. Defaults to BASIC.
    """
    BASIC = "BASIC"
    """
    The basic view omits fields which can be large or can contain sensitive data. This view does not include the body in AppEngineHttpRequest. Bodies are desirable to return only when needed, because they can be large and because of the sensitivity of the data that you choose to store in it.
    """
    FULL = "FULL"
    """
    All information is returned. Authorization for FULL requires `cloudtasks.tasks.fullView` [Google IAM](https://cloud.google.com/iam/) permission on the Queue resource.
    """


class UriOverrideScheme(str, Enum):
    """
    Scheme override. When specified, the task URI scheme is replaced by the provided value (HTTP or HTTPS).
    """
    SCHEME_UNSPECIFIED = "SCHEME_UNSPECIFIED"
    """
    Scheme unspecified. Defaults to HTTPS.
    """
    HTTP = "HTTP"
    """
    Convert the scheme to HTTP, e.g., https://www.google.ca will change to http://www.google.ca.
    """
    HTTPS = "HTTPS"
    """
    Convert the scheme to HTTPS, e.g., http://www.google.ca will change to https://www.google.ca.
    """


class UriOverrideUriOverrideEnforceMode(str, Enum):
    """
    URI Override Enforce Mode When specified, determines the Target UriOverride mode. If not specified, it defaults to ALWAYS.
    """
    URI_OVERRIDE_ENFORCE_MODE_UNSPECIFIED = "URI_OVERRIDE_ENFORCE_MODE_UNSPECIFIED"
    """
    OverrideMode Unspecified. Defaults to ALWAYS.
    """
    IF_NOT_EXISTS = "IF_NOT_EXISTS"
    """
    In the IF_NOT_EXISTS mode, queue-level configuration is only applied where task-level configuration does not exist.
    """
    ALWAYS = "ALWAYS"
    """
    In the ALWAYS mode, queue-level configuration overrides all task-level configuration
    """
