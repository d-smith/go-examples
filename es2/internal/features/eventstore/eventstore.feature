@eventstore
Feature: Event Store

    Scenario: Events for an Aggregate
        Given an aggregate ID
        And an event store
        When the aggregate has uncommitted events
        And the events are stored
        Then the events for that aggregate can be retrieved

    Scenario: Events for a specific aggregate id can be retrieved
        Given two aggregates
        And an event store containing all events for both aggregates
        When the events for an aggregate are retrieved
        Then only the events assocaited with the specific aggregate are retrieved
        And the events are ordered