// This file was generated by typhen-api

using System.Collections.Generic;

namespace TyphenApi.Type.Submarine
{
    public partial class User : TyphenApi.TypeBase
    {
        protected static readonly SerializationInfo<User, long> id = new SerializationInfo<User, long>("id", false, (x) => x.Id, (x, v) => x.Id = v);
        public long Id { get; set; }
        protected static readonly SerializationInfo<User, string> name = new SerializationInfo<User, string>("name", false, (x) => x.Name, (x, v) => x.Name = v);
        public string Name { get; set; }
    }
}
