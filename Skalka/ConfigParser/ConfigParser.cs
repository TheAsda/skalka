using System.Runtime.CompilerServices;
using Skalka.IO;
using YamlDotNet.Serialization;
using YamlDotNet.Serialization.NamingConventions;

namespace Skalka.ConfigParser
{
    public class ConfigParser: IConfigParser
    {
        private readonly IPathReader _reader;
        private readonly IDeserializer _deserializer;

        public ConfigParser(IPathReader reader)
        {
            _reader = reader;
            _deserializer = new DeserializerBuilder().WithNamingConvention(CamelCaseNamingConvention.Instance).Build();
        }

        public T ReadConfig<T>(string path)
        {
            var data = _reader.Read(path);
            return _deserializer.Deserialize<T>(data);
        }
    }
}