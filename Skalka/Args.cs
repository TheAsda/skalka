using CommandLine;

namespace Skalka
{
    public record Args
    {
        [Option('c', "config", HelpText = "Path to YAML flow config file")]
        public string ConfigFile { get; set; }
    }
}