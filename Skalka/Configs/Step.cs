using System.Collections.Generic;

namespace Skalka.Configs
{
    public enum StepType
    {
        Step,
        StepWithOptions,
        StepPlugin
    }

    public record Step
    {
        public string Name { get; set; }
        public List<string> Env { get; set; }
        public string Run { get; set; }
        public string Dir { get; set; }
        public StepType Type { get; set; }
        public string VariableName { get; set; }
        public Dictionary<string, StepOption> Options { get; set; }
    }
}