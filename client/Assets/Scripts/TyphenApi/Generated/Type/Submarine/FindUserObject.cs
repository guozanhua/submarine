// This file was generated by typhen-api

using System.Collections.Generic;

namespace TyphenApi.Type.Submarine
{
    public partial class FindUserObject : TyphenApi.TypeBase
    {
        protected static readonly SerializationInfo<FindUserObject, TyphenApi.Type.Submarine.User> user = new SerializationInfo<FindUserObject, TyphenApi.Type.Submarine.User>("user", true, (x) => x.User, (x, v) => x.User = v);
        public TyphenApi.Type.Submarine.User User { get; set; }
    }
}
