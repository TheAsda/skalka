namespace Skalka.Queue
{
    public interface IQueueProcessorFactory
    {
        IQueueProcessor Create();
    }
}