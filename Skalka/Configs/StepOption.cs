namespace Skalka.Configs
{
    public record StepOption
    {
        public string Name { get; set; }
        public bool IsDefault { get; set; }
        public string Run { get; set; }
        public bool Skip { get; set; }
    }
}