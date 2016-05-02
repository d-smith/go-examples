@aggregate
Feature: Event Sourced Aggregate

    Scenario: Aggregate created
        Given an event sourced aggregate
        When I create new instances
        Then the instances have unique IDs