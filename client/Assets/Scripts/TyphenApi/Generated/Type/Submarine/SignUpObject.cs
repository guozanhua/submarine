// This file was generated by typhen-api

using System.Collections.Generic;

namespace TyphenApi.Type.Submarine
{
    public partial class SignUpObject : TyphenApi.TypeBase
    {
        protected static readonly SerializationInfo<SignUpObject, TyphenApi.Type.Submarine.User> user = new SerializationInfo<SignUpObject, TyphenApi.Type.Submarine.User>("user", false, (x) => x.User, (x, v) => x.User = v);
        public TyphenApi.Type.Submarine.User User { get; set; }
    }
}
