// This file was generated by typhen-api

using System.Collections.Generic;

namespace TyphenApi.Type.Submarine.Configuration
{
    public partial class Client : TyphenApi.TypeBase<Client>
    {
        protected static readonly SerializationInfo<Client, string> version = new SerializationInfo<Client, string>("version", false, (x) => x.Version, (x, v) => x.Version = v);
        public string Version { get; set; }
        protected static readonly SerializationInfo<Client, string> apiServerBaseUri = new SerializationInfo<Client, string>("api_server_base_uri", false, (x) => x.ApiServerBaseUri, (x, v) => x.ApiServerBaseUri = v);
        public string ApiServerBaseUri { get; set; }
    }
}
