﻿using UnityEngine;
using System;
using System.Collections;

namespace TyphenApi
{
    public class WebApiRequestSenderWWW : IWebApiRequestSender
    {
        public IEnumerator Send(IWebApiRequest request)
        {
            switch (request.Method)
            {
                case HttpMethod.Get:
                {
                    var www = new WWW(request.Uri + request.Body.ToQueryString(), null, request.Headers);
                    yield return www;
                    request.OnSendComplete(www.responseHeaders, www.bytes, www.error);
                    break;
                }
                case HttpMethod.Post:
                {
                    var www = new WWW(request.Uri.ToString(), request.BodyBytes, request.Headers);
                    yield return www;
                    request.OnSendComplete(www.responseHeaders, www.bytes, www.error);
                    break;
                }
                default:
                {
                    throw new NotImplementedException("Put or Delete HTTPMethod is not implemented yet");
                }
            }
        }
    }
}