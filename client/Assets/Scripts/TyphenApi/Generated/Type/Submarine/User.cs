// This file was generated by typhen-api

using System.Collections.Generic;

namespace TyphenApi.Type.Submarine
{
    [MessagePack.MessagePackObject]
    [Newtonsoft.Json.JsonObject(Newtonsoft.Json.MemberSerialization.OptIn)]
    public partial class User : TyphenApi.TypeBase<User>
    {
        [TyphenApi.QueryStringProperty("id", false)]
        [MessagePack.Key("id")]
        [Newtonsoft.Json.JsonProperty("id")]
        [Newtonsoft.Json.JsonRequired]
        public long Id { get; set; }
        [TyphenApi.QueryStringProperty("name", false)]
        [MessagePack.Key("name")]
        [Newtonsoft.Json.JsonProperty("name")]
        [Newtonsoft.Json.JsonRequired]
        public string Name { get; set; }
    }
}
