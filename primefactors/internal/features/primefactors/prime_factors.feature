Feature: Primefactors

  @primefactors
  Scenario:
    Given A prime factor resource value of 8125
    When I call the prime factors service
    Then The prime factors for the resouce value are returned