from unittest import TestCase

import pytest

from polyaxon.automl.search_managers.hyperband.manager import HyperbandManager
from polyaxon.schemas.polyflow.workflows import HyperbandConfig


@pytest.mark.automl_mark
class TestHyperbandSearchManager(TestCase):
    DISABLE_RUNNER = True
    DISABLE_EXECUTOR = True
    DISABLE_AUDITOR = True

    def setUp(self):
        super().setUp()
        config = HyperbandConfig.from_dict(
            {
                "concurrency": 2,
                "max_iter": 10,
                "eta": 3,
                "resource": {"name": "steps", "type": "float"},
                "resume": False,
                "metric": {"name": "loss", "optimization": "minimize"},
                "matrix": {
                    "feature1": {"kind": "choice", "value": [1, 2, 3]},
                    "feature2": {"kind": "linspace", "value": [1, 2, 5]},
                    "feature3": {"kind": "range", "value": [1, 5, 1]},
                },
            }
        )
        self.manager1 = HyperbandManager(config=config)

        config = HyperbandConfig.from_dict(
            {
                "concurrency": 2,
                "max_iter": 81,
                "eta": 3,
                "resource": {"name": "size", "type": "int"},
                "resume": False,
                "metric": {"name": "loss", "optimization": "minimize"},
                "matrix": {
                    "feature1": {"kind": "choice", "value": [1, 2, 3]},
                    "feature2": {"kind": "linspace", "value": [1, 2, 5]},
                    "feature3": {"kind": "range", "value": [1, 5, 1]},
                    "feature4": {"kind": "range", "value": [1, 5, 1]},
                },
            }
        )
        self.manager2 = HyperbandManager(config=config)

    def test_hyperband_config(self):
        assert HyperbandManager.CONFIG == HyperbandConfig

    @staticmethod
    def almost_equal(value, value_compare):
        assert value_compare - 0.02 <= value <= value_compare + 0.02

    def test_hyperband_properties(self):
        # Manager1
        assert self.manager1.max_iter == 10
        assert self.manager1.eta == 3
        assert self.manager1.s_max == 2
        assert self.manager1.B == (self.manager1.s_max + 1) * self.manager1.max_iter

        assert self.manager2.max_iter == 81
        assert self.manager2.eta == 3
        assert self.manager2.s_max == 4
        assert self.manager2.B == (self.manager2.s_max + 1) * self.manager2.max_iter

    def test_get_bracket(self):
        # Manager1
        assert self.manager1.get_bracket(iteration=0) == 2
        assert self.manager1.get_bracket(iteration=1) == 1
        assert self.manager1.get_bracket(iteration=2) == 0

        # Manager2
        assert self.manager2.get_bracket(iteration=0) == 4
        assert self.manager2.get_bracket(iteration=1) == 3
        assert self.manager2.get_bracket(iteration=2) == 2
        assert self.manager2.get_bracket(iteration=3) == 1
        assert self.manager2.get_bracket(iteration=4) == 0

    def test_get_n_runs(self):
        # Manager1
        assert self.manager1.get_n_runs(bracket=2) == 9
        assert self.manager1.get_n_runs(bracket=1) == 5
        assert self.manager1.get_n_runs(bracket=0) == 3

        # Manager2
        assert self.manager2.get_n_runs(bracket=4) == 81
        assert self.manager2.get_n_runs(bracket=3) == 34
        assert self.manager2.get_n_runs(bracket=2) == 15
        assert self.manager2.get_n_runs(bracket=1) == 8
        assert self.manager2.get_n_runs(bracket=0) == 5

    def test_get_n_runs_to_keep(self):
        # Number of config to keep [/n_suggestions, /iteration]

        # Manager1
        # Iteration == 0
        assert self.manager1.get_n_runs_to_keep(n_runs=9, bracket_iteration=0) == 3
        assert self.manager1.get_n_runs_to_keep(n_runs=9, bracket_iteration=1) == 1
        assert self.manager1.get_n_runs_to_keep(n_runs=9, bracket_iteration=2) == 0

        assert (
            self.manager1.get_n_runs_to_keep_for_iteration(
                iteration=0, bracket_iteration=0
            )
            == 3
        )
        assert (
            self.manager1.get_n_runs_to_keep_for_iteration(
                iteration=0, bracket_iteration=1
            )
            == 1
        )
        assert (
            self.manager1.get_n_runs_to_keep_for_iteration(
                iteration=0, bracket_iteration=2
            )
            == 0
        )

        # Iteration == 1
        assert self.manager1.get_n_runs_to_keep(n_runs=5, bracket_iteration=0) == 1
        assert self.manager1.get_n_runs_to_keep(n_runs=5, bracket_iteration=1) == 0

        assert (
            self.manager1.get_n_runs_to_keep_for_iteration(
                iteration=1, bracket_iteration=0
            )
            == 1
        )
        assert (
            self.manager1.get_n_runs_to_keep_for_iteration(
                iteration=1, bracket_iteration=1
            )
            == 0
        )

        # Iteration == 2
        assert self.manager1.get_n_runs_to_keep(n_runs=3, bracket_iteration=0) == 1
        assert (
            self.manager1.get_n_runs_to_keep_for_iteration(
                iteration=2, bracket_iteration=0
            )
            == 1
        )

        # Manager2
        # Iteration == 0
        assert self.manager2.get_n_runs_to_keep(n_runs=81, bracket_iteration=0) == 27
        assert self.manager2.get_n_runs_to_keep(n_runs=81, bracket_iteration=1) == 9
        assert self.manager2.get_n_runs_to_keep(n_runs=81, bracket_iteration=2) == 3
        assert self.manager2.get_n_runs_to_keep(n_runs=81, bracket_iteration=3) == 1
        assert self.manager2.get_n_runs_to_keep(n_runs=81, bracket_iteration=4) == 0

        assert (
            self.manager2.get_n_runs_to_keep_for_iteration(
                iteration=0, bracket_iteration=0
            )
            == 27
        )
        assert (
            self.manager2.get_n_runs_to_keep_for_iteration(
                iteration=0, bracket_iteration=1
            )
            == 9
        )
        assert (
            self.manager2.get_n_runs_to_keep_for_iteration(
                iteration=0, bracket_iteration=2
            )
            == 3
        )
        assert (
            self.manager2.get_n_runs_to_keep_for_iteration(
                iteration=0, bracket_iteration=3
            )
            == 1
        )
        assert (
            self.manager2.get_n_runs_to_keep_for_iteration(
                iteration=0, bracket_iteration=4
            )
            == 0
        )

        # Iteration == 1
        assert self.manager2.get_n_runs_to_keep(n_runs=34, bracket_iteration=0) == 11
        assert self.manager2.get_n_runs_to_keep(n_runs=34, bracket_iteration=1) == 3
        assert self.manager2.get_n_runs_to_keep(n_runs=34, bracket_iteration=2) == 1
        assert self.manager2.get_n_runs_to_keep(n_runs=34, bracket_iteration=3) == 0

        assert (
            self.manager2.get_n_runs_to_keep_for_iteration(
                iteration=1, bracket_iteration=0
            )
            == 11
        )
        assert (
            self.manager2.get_n_runs_to_keep_for_iteration(
                iteration=1, bracket_iteration=1
            )
            == 3
        )
        assert (
            self.manager2.get_n_runs_to_keep_for_iteration(
                iteration=1, bracket_iteration=2
            )
            == 1
        )
        assert (
            self.manager2.get_n_runs_to_keep_for_iteration(
                iteration=1, bracket_iteration=3
            )
            == 0
        )

        # Iteration == 2
        assert self.manager2.get_n_runs_to_keep(n_runs=15, bracket_iteration=0) == 5
        assert self.manager2.get_n_runs_to_keep(n_runs=15, bracket_iteration=1) == 1
        assert self.manager2.get_n_runs_to_keep(n_runs=15, bracket_iteration=2) == 0

        assert (
            self.manager2.get_n_runs_to_keep_for_iteration(
                iteration=2, bracket_iteration=0
            )
            == 5
        )
        assert (
            self.manager2.get_n_runs_to_keep_for_iteration(
                iteration=2, bracket_iteration=1
            )
            == 1
        )
        assert (
            self.manager2.get_n_runs_to_keep_for_iteration(
                iteration=2, bracket_iteration=2
            )
            == 0
        )

        # Iteration == 3
        assert self.manager2.get_n_runs_to_keep(n_runs=8, bracket_iteration=0) == 2
        assert self.manager2.get_n_runs_to_keep(n_runs=8, bracket_iteration=1) == 0

        assert (
            self.manager2.get_n_runs_to_keep_for_iteration(
                iteration=3, bracket_iteration=0
            )
            == 2
        )
        assert (
            self.manager2.get_n_runs_to_keep_for_iteration(
                iteration=3, bracket_iteration=1
            )
            == 0
        )

        # Iteration == 4
        assert self.manager2.get_n_runs_to_keep(n_runs=5, bracket_iteration=0) == 1

        assert (
            self.manager2.get_n_runs_to_keep_for_iteration(
                iteration=4, bracket_iteration=0
            )
            == 1
        )

    def test_get_resources(self):
        # Number of resources [/bracket, /iteration]

        # Manager1
        self.almost_equal(self.manager1.get_resources(bracket=2), 1.11)
        self.almost_equal(self.manager1.get_resources(bracket=1), 3.33)
        self.almost_equal(self.manager1.get_resources(bracket=0), 10)

        self.almost_equal(self.manager1.get_resources_for_iteration(iteration=0), 1.11)
        self.almost_equal(self.manager1.get_resources_for_iteration(iteration=1), 3.33)
        self.almost_equal(self.manager1.get_resources_for_iteration(iteration=2), 10)

        # Manager2
        self.almost_equal(self.manager2.get_resources(bracket=4), 1)
        self.almost_equal(self.manager2.get_resources(bracket=3), 3)
        self.almost_equal(self.manager2.get_resources(bracket=2), 9)
        self.almost_equal(self.manager2.get_resources(bracket=1), 27)
        self.almost_equal(self.manager2.get_resources(bracket=0), 81)

        self.almost_equal(self.manager2.get_resources_for_iteration(iteration=0), 1)
        self.almost_equal(self.manager2.get_resources_for_iteration(iteration=1), 3)
        self.almost_equal(self.manager2.get_resources_for_iteration(iteration=2), 9)
        self.almost_equal(self.manager2.get_resources_for_iteration(iteration=3), 27)
        self.almost_equal(self.manager2.get_resources_for_iteration(iteration=4), 81)

    def test_get_n_resources(self):
        # Number of iteration resources

        # Manager1
        # Iteration == 0
        self.almost_equal(
            self.manager1.get_n_resources(n_resources=1.11, bracket_iteration=0), 1.11
        )
        self.almost_equal(
            self.manager1.get_n_resources(n_resources=1.11, bracket_iteration=1), 3.33
        )
        self.almost_equal(
            self.manager1.get_n_resources(n_resources=1.11, bracket_iteration=2), 9.99
        )

        self.almost_equal(
            self.manager1.get_n_resources_for_iteration(
                iteration=0, bracket_iteration=0
            ),
            1.11,
        )
        self.almost_equal(
            self.manager1.get_n_resources_for_iteration(
                iteration=0, bracket_iteration=1
            ),
            3.33,
        )
        self.almost_equal(
            self.manager1.get_n_resources_for_iteration(
                iteration=0, bracket_iteration=2
            ),
            9.99,
        )

        # Iteration == 1
        self.almost_equal(
            self.manager1.get_n_resources(n_resources=3.33, bracket_iteration=0), 3.33
        )
        self.almost_equal(
            self.manager1.get_n_resources(n_resources=3.33, bracket_iteration=1), 9.99
        )

        self.almost_equal(
            self.manager1.get_n_resources_for_iteration(
                iteration=1, bracket_iteration=0
            ),
            3.33,
        )
        self.almost_equal(
            self.manager1.get_n_resources_for_iteration(
                iteration=1, bracket_iteration=1
            ),
            9.99,
        )

        # Iteration == 2
        self.almost_equal(
            self.manager1.get_n_resources(n_resources=9.99, bracket_iteration=0), 9.99
        )

        self.almost_equal(
            self.manager1.get_n_resources_for_iteration(
                iteration=2, bracket_iteration=0
            ),
            9.99,
        )

        # Manager2
        # Iteration == 0
        self.almost_equal(
            self.manager2.get_n_resources(n_resources=1, bracket_iteration=0), 1
        )
        self.almost_equal(
            self.manager2.get_n_resources(n_resources=1, bracket_iteration=1), 3
        )
        self.almost_equal(
            self.manager2.get_n_resources(n_resources=1, bracket_iteration=2), 9
        )
        self.almost_equal(
            self.manager2.get_n_resources(n_resources=1, bracket_iteration=3), 27
        )
        self.almost_equal(
            self.manager2.get_n_resources(n_resources=1, bracket_iteration=4), 81
        )

        self.almost_equal(
            self.manager2.get_n_resources_for_iteration(
                iteration=0, bracket_iteration=0
            ),
            1,
        )
        self.almost_equal(
            self.manager2.get_n_resources_for_iteration(
                iteration=0, bracket_iteration=1
            ),
            3,
        )
        self.almost_equal(
            self.manager2.get_n_resources_for_iteration(
                iteration=0, bracket_iteration=2
            ),
            9,
        )
        self.almost_equal(
            self.manager2.get_n_resources_for_iteration(
                iteration=0, bracket_iteration=3
            ),
            27,
        )
        self.almost_equal(
            self.manager2.get_n_resources_for_iteration(
                iteration=0, bracket_iteration=4
            ),
            81,
        )

        # Iteration == 1
        self.almost_equal(
            self.manager2.get_n_resources(n_resources=3, bracket_iteration=0), 3
        )
        self.almost_equal(
            self.manager2.get_n_resources(n_resources=3, bracket_iteration=1), 9
        )
        self.almost_equal(
            self.manager2.get_n_resources(n_resources=3, bracket_iteration=2), 27
        )
        self.almost_equal(
            self.manager2.get_n_resources(n_resources=3, bracket_iteration=3), 81
        )

        self.almost_equal(
            self.manager2.get_n_resources_for_iteration(
                iteration=1, bracket_iteration=0
            ),
            3,
        )
        self.almost_equal(
            self.manager2.get_n_resources_for_iteration(
                iteration=1, bracket_iteration=1
            ),
            9,
        )
        self.almost_equal(
            self.manager2.get_n_resources_for_iteration(
                iteration=1, bracket_iteration=2
            ),
            27,
        )
        self.almost_equal(
            self.manager2.get_n_resources_for_iteration(
                iteration=1, bracket_iteration=3
            ),
            81,
        )

        # Iteration == 2
        self.almost_equal(
            self.manager2.get_n_resources(n_resources=9, bracket_iteration=0), 9
        )
        self.almost_equal(
            self.manager2.get_n_resources(n_resources=9, bracket_iteration=1), 27
        )
        self.almost_equal(
            self.manager2.get_n_resources(n_resources=9, bracket_iteration=2), 81
        )

        self.almost_equal(
            self.manager2.get_n_resources_for_iteration(
                iteration=2, bracket_iteration=0
            ),
            9,
        )
        self.almost_equal(
            self.manager2.get_n_resources_for_iteration(
                iteration=2, bracket_iteration=1
            ),
            27,
        )
        self.almost_equal(
            self.manager2.get_n_resources_for_iteration(
                iteration=2, bracket_iteration=2
            ),
            81,
        )

        # Iteration == 3
        self.almost_equal(
            self.manager2.get_n_resources(n_resources=27, bracket_iteration=0), 27
        )
        self.almost_equal(
            self.manager2.get_n_resources(n_resources=27, bracket_iteration=1), 81
        )

        self.almost_equal(
            self.manager2.get_n_resources_for_iteration(
                iteration=3, bracket_iteration=0
            ),
            27,
        )
        self.almost_equal(
            self.manager2.get_n_resources_for_iteration(
                iteration=3, bracket_iteration=1
            ),
            81,
        )

        # Iteration == 4
        self.almost_equal(
            self.manager2.get_n_resources(n_resources=81, bracket_iteration=0), 81
        )

        self.almost_equal(
            self.manager2.get_n_resources_for_iteration(
                iteration=4, bracket_iteration=0
            ),
            81,
        )

    def test_should_reschedule(self):
        assert (
            self.manager1.should_reschedule(iteration=0, bracket_iteration=0) is False
        )
        assert (
            self.manager1.should_reschedule(iteration=0, bracket_iteration=1) is False
        )
        assert self.manager1.should_reschedule(iteration=0, bracket_iteration=2) is True
        assert self.manager1.should_reschedule(iteration=0, bracket_iteration=3) is True
        assert (
            self.manager1.should_reschedule(iteration=1, bracket_iteration=0) is False
        )
        assert self.manager1.should_reschedule(iteration=1, bracket_iteration=1) is True
        assert self.manager1.should_reschedule(iteration=1, bracket_iteration=2) is True
        assert (
            self.manager1.should_reschedule(iteration=2, bracket_iteration=0) is False
        )
        assert (
            self.manager1.should_reschedule(iteration=2, bracket_iteration=1) is False
        )
        assert (
            self.manager1.should_reschedule(iteration=5, bracket_iteration=0) is False
        )

    def test_should_reduce_configs(self):
        assert (
            self.manager1.should_reduce_configs(iteration=0, bracket_iteration=0)
            is True
        )
        assert (
            self.manager1.should_reduce_configs(iteration=0, bracket_iteration=1)
            is True
        )
        assert (
            self.manager1.should_reduce_configs(iteration=0, bracket_iteration=2)
            is False
        )
        assert (
            self.manager1.should_reduce_configs(iteration=0, bracket_iteration=3)
            is False
        )
        assert (
            self.manager1.should_reduce_configs(iteration=1, bracket_iteration=0)
            is True
        )
        assert (
            self.manager1.should_reduce_configs(iteration=1, bracket_iteration=1)
            is False
        )
        assert (
            self.manager1.should_reduce_configs(iteration=1, bracket_iteration=2)
            is False
        )
        assert (
            self.manager1.should_reduce_configs(iteration=2, bracket_iteration=0)
            is True
        )
        assert (
            self.manager1.should_reduce_configs(iteration=2, bracket_iteration=1)
            is False
        )
        assert (
            self.manager1.should_reduce_configs(iteration=5, bracket_iteration=0)
            is False
        )

    def test_get_suggestions(self):
        suggestions = self.manager1.get_suggestions(iteration=0, bracket_iteration=2)

        assert len(suggestions) == 9
        for suggestion in suggestions:
            assert "steps" in suggestion
            self.almost_equal(suggestion["steps"], 9.99)
            assert "feature1" in suggestion
            assert "feature2" in suggestion
            assert "feature3" in suggestion

        # Fake iteration
        suggestions = self.manager1.get_suggestions(iteration=1, bracket_iteration=0)
        assert len(suggestions) == 5
        for suggestion in suggestions:
            assert "steps" in suggestion
            self.almost_equal(suggestion["steps"], 3.33)
            assert "feature1" in suggestion
            assert "feature2" in suggestion
            assert "feature3" in suggestion

        # Fake iteration
        suggestions = self.manager1.get_suggestions(iteration=2, bracket_iteration=0)
        assert len(suggestions) == 3
        for suggestion in suggestions:
            assert "steps" in suggestion
            self.almost_equal(suggestion["steps"], 9.99)
            assert "feature1" in suggestion
            assert "feature2" in suggestion
            assert "feature3" in suggestion

        # Manager2
        # Fake iteration
        suggestions = self.manager2.get_suggestions(iteration=2, bracket_iteration=0)
        assert len(suggestions) == 15
        for suggestion in suggestions:
            assert "size" in suggestion
            self.almost_equal(suggestion["size"], 9)
            assert "feature1" in suggestion
            assert "feature2" in suggestion
            assert "feature3" in suggestion
            assert "feature4" in suggestion

        # Fake iteration
        suggestions = self.manager2.get_suggestions(iteration=4, bracket_iteration=0)
        assert len(suggestions) == 5
        for suggestion in suggestions:
            assert "size" in suggestion
            self.almost_equal(suggestion["size"], 81)
            assert "feature1" in suggestion
            assert "feature2" in suggestion
            assert "feature3" in suggestion
            assert "feature4" in suggestion
