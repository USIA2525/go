package  animals

import (
    "testing"
)

func TestElephantFeed(t *testing.T) {
    expect := "Grass"
    actual := ElephantFeed()


    if expect != actual {
        t.Errorf ("%s != %s", expect, actual)
    }
}

func TestMonkeyFeed(t *testing.T) {
    expect := "Banna"
    actual := MonkeyFeed()

    if expect != actual {
        t.Errorf ("%s != %s", expect, actual)
    }
}

func TestRabbitFeed(t *testing.T) {
    expect := "Carrot"
    actual := RabbitFeed()

    if expect != actual {
        t.Errorf ("%s != %s", expect, actual)
    }
}
