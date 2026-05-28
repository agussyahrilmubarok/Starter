using System.Text.Json.Serialization;

namespace Backend.Application.DTOs;

public class ApiResponse<T>
{
    [JsonPropertyName("message")] public string Message { get; set; }

    [JsonPropertyName("data")]
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
    public T? Data { get; set; }

    [JsonPropertyName("errors")]
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
    public Dictionary<string, string>? Errors { get; set; }

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