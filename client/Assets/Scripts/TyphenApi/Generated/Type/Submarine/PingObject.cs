// This file was generated by typhen-api

using System.Collections.Generic;

namespace TyphenApi.Type.Submarine
{
    public partial class PingObject : TyphenApi.TypeBase
    {
        protected static readonly SerializationInfo<PingObject, string> message = new SerializationInfo<PingObject, string>("message", false, (x) => x.Message, (x, v) => x.Message = v);
        public string Message { get; set; }
    }
}
