using System.IO;
using Skalka.IO.Exceptions;

namespace Skalka.IO
{
    public class PathReader : IPathReader
    {
        public string Read(string path, bool relative = true)
        {
            var fullPath = relative ? Path.Combine(Directory.GetCurrentDirectory(), path) : path;
            fullPath = Path.GetFullPath(fullPath);
            if (!File.Exists(fullPath))
            {
                throw new NotExistException(fullPath);
            }

            return File.ReadAllText(fullPath);
        }
    }
}