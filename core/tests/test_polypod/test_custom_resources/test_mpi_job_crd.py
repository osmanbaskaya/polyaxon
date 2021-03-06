#!/usr/bin/python
#
# Copyright 2018-2020 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
from tests.test_polypod.test_custom_resources.base_kubeflow import (
    BaseKubeflowCRDTestCase,
)

from polyaxon.k8s.custom_resources.crd import get_custom_object
from polyaxon.polyflow import V1Notification, V1NotificationTrigger
from polyaxon.polyflow.environment import V1Environment
from polyaxon.polyflow.termination import V1Termination
from polyaxon.polypod.custom_resources import get_mpi_job_custom_resource


class TestPytorchJobCRD(BaseKubeflowCRDTestCase):
    def test_get_pytorch_job_custom_resource_with_no_workers(self):
        termination = V1Termination(max_retries=5, ttl=10, timeout=10)
        environment = V1Environment(
            labels={"foo": "bar"},
            annotations={"foo": "bar"},
            node_selector={"foo": "bar"},
            node_name="foo",
            restart_policy="never",
        )
        custom_object = {
            "mpiJobSpec": {"cleanPodPolicy": "All"},
            "termination": {
                "backoffLimit": termination.max_retries,
                "activeDeadlineSeconds": termination.timeout,
                "ttlSecondsAfterFinished": termination.ttl,
            },
            "collectLogs": False,
            "syncStatuses": False,
            "notifications": [],
        }
        expected_crd = get_custom_object(
            namespace="default",
            resource_name="foo",
            kind="Operation",
            api_version="core.polyaxon.com/v1",
            labels={"foo": "bar"},
            custom_object=custom_object,
        )

        crd = get_mpi_job_custom_resource(
            namespace="default",
            resource_name="foo",
            launcher=None,
            worker=None,
            slots_per_worker=None,
            clean_pod_policy=None,
            termination=termination,
            collect_logs=False,
            sync_statuses=False,
            notifications=None,
            labels=environment.labels,
        )

        assert crd == expected_crd

    def test_get_job_custom_resource(self):
        termination = V1Termination(max_retries=5, ttl=10, timeout=10)
        environment = V1Environment(
            labels={"foo": "bar"},
            annotations={"foo": "bar"},
            node_selector={"foo": "bar"},
            node_name="foo",
            restart_policy="never",
        )
        notifications = [
            V1Notification(connection="test", trigger=V1NotificationTrigger.DONE)
        ]
        launcher, launcher_replica_template = self.get_replica(environment)
        worker, worker_replica_template = self.get_replica(environment)
        template_spec = {
            "cleanPodPolicy": "Running",
            "slotsPerWorker": 12,
            "Launcher": launcher_replica_template,
            "Worker": worker_replica_template,
        }
        custom_object = {
            "mpiJobSpec": template_spec,
            "termination": {
                "backoffLimit": termination.max_retries,
                "activeDeadlineSeconds": termination.timeout,
                "ttlSecondsAfterFinished": termination.ttl,
            },
            "collectLogs": True,
            "syncStatuses": True,
            "notifications": [n.to_operator_notation() for n in notifications],
        }

        expected_crd = get_custom_object(
            namespace="default",
            resource_name="foo",
            kind="Operation",
            api_version="core.polyaxon.com/v1",
            labels={"foo": "bar"},
            custom_object=custom_object,
        )

        crd = get_mpi_job_custom_resource(
            namespace="default",
            resource_name="foo",
            launcher=launcher,
            worker=worker,
            slots_per_worker=12,
            clean_pod_policy="Running",
            termination=termination,
            collect_logs=True,
            sync_statuses=True,
            notifications=notifications,
            labels=environment.labels,
        )

        assert crd == expected_crd
