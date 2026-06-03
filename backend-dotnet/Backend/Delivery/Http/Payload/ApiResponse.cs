using System.Text.Json.Serialization;

namespace Backend.Delivery.Http.Payload;

public class ApiResponse<T>
{
    [JsonPropertyName("message")]
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
    public string Message { get; set; } = null!;

    [JsonPropertyName("data")]
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
    public T? Data { get; set; }

    [JsonPropertyName("errors")]
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
    public Dictionary<string, string>? Errors { get; set; }

    [JsonConstructor]
    public ApiResponse() { }

    public ApiResponse(string message, T data)
    {
        Message = message;
        Data = data;
    }

    public ApiResponse(string message, Dictionary<string, string> errors)
    {
        Message = message;
        Errors = errors;
    }
}