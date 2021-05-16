using System;

namespace Skalka.IO.Exceptions
{
    public class NotExistException : Exception
    {
        public NotExistException(string fullPath) : base($"File does not exist '{fullPath}'")
        {
        }
    }
}