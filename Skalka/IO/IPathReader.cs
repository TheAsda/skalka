namespace Skalka.IO
{
    public interface IPathReader
    {
        public string Read(string path, bool relative = true);
    }
}