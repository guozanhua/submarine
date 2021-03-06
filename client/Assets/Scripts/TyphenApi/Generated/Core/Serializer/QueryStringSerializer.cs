﻿// This file was generated by typhen-api

using System;
using System.Collections.Generic;
using System.Text;
using System.Linq;
using System.Reflection;

namespace TyphenApi
{
    [AttributeUsage(AttributeTargets.Property, Inherited = false)]
    public class QueryStringPropertyAttribute : Attribute
    {
        public readonly string Name;
        public readonly bool IsOptional;

        public QueryStringPropertyAttribute(string name, bool isOptional)
        {
            Name = name;
            IsOptional = isOptional;
        }
    }

    public class QueryStringSerializer : ISerializer
    {
        public byte[] Serialize<T>(T obj)
        {
            var texts = new List<string>();

            foreach (var prop in obj.GetType().GetProperties())
            {
                foreach (var attr in prop.GetCustomAttributes(true).OfType<QueryStringPropertyAttribute>())
                {
                    var value = prop.GetValue(obj, null);
                    var valueType = prop.PropertyType;

                    if (value == null)
                    {
                        if (attr.IsOptional)
                        {
                            continue;
                        }
                        else
                        {
                            var message = string.Format("{0}.{1} is not allowed to be null.", obj.GetType().FullName, attr.Name);
                            throw new SerializationFailureException(message);
                        }
                    }

                    if (IsSerializableValue(value, valueType))
                    {
                        var fixedValue = valueType.IsEnum ? (int)value : value;
                        var keyValueText = string.Format("{0}={1}",
                            Uri.EscapeDataString(attr.Name),
                            Uri.EscapeDataString(fixedValue.ToString())
                        );
                        texts.Add(keyValueText);
                    }
                    else
                    {
                        var message = string.Format("Failed to serialize {0} ({1}) to {2}.{3}", value, valueType, obj.GetType().FullName, attr.Name);
                        throw new SerializationFailureException(message);
                    }
                }
            }

            return texts.Count > 0 ?
                Encoding.ASCII.GetBytes("?" + string.Join("&", texts.ToArray())) :
                Encoding.ASCII.GetBytes(string.Empty);
        }

        bool IsSerializableValue(object value, System.Type valueType)
        {
            return valueType.IsPrimitive || valueType.IsEnum || value is string;
        }
    }
}
