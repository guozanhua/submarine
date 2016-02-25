﻿using UnityEngine;
using DG.Tweening;

namespace Submarine.Battle
{
    public class SubmarineView : MonoBehaviour
    {
        [SerializeField]
        GameObject model;
        [SerializeField]
        Transform torpedoLaunchSite;
        [SerializeField]
        Transform decoyLaunchSite;
        [SerializeField]
        Transform lookoutLaunchSite;
        [SerializeField]
        GameObject streamEffect;
        [SerializeField]
        Material enemySubmarineMaterial;

        public Vector3 TorpedoLaunchSitePosition { get { return torpedoLaunchSite.position; } }
        public Vector3 DecoyLaunchSitePosition { get { return decoyLaunchSite.position; } }
        public Quaternion DecoyLaunchSiteRotation { get { return decoyLaunchSite.rotation; } }
        public Vector3 LookoutLaunchSitePosition { get { return lookoutLaunchSite.position; } }

        bool isMine;
        public bool IsMine
        {
            get { return isMine; }
            set
            {
                isMine = value;
                if (!IsMine)
                {
                    model.GetComponent<MeshRenderer>().material = enemySubmarineMaterial;
                }
            }
        }

        void Start()
        {
            model.transform.DOLocalMoveY(-0.25f, 3f)
                .SetEase(Ease.InOutQuad)
                .SetLoops(-1, LoopType.Yoyo);
        }
    }
}