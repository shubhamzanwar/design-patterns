# Adapter Pattern 🔌

The Adapter pattern is a structural design pattern. It allows 2 or more incompatible objects to interface with each other. The first thing that comes to mind it probably the adapters for wall-sockets that allow your indian phone charger to fit into european wall sockets! 😂. That's exactly what adapters are. They are wrapper classes written around non-uniform classes so that the end result exposes a uniform API.

In case this is difficult to grasp, the following example should help:

### Translator

Consider that you're a regular english speaking person who has just migrated to Spain. 🇪🇸
Your neighbor is a pleasant looking person and you'd like to be friends with this individual. However, there's one problem. 😨

_Every one in your neighborhood speaks Español exclusively. Oh, also, you don't speak spanish_ 😞

In order to communicate with your potential friend, you would need some sort of translator to help convert your speech from English to Spanish. In programming terms, this translator is what we call an adapter!

Let's understand at how this would look in code.

As a first step, let's create the interface that a normal spanish-speaking person would implement along with the structs for both, the spanish and the english speaking person.

```go
type SpeakSpanish interface {
    greetInSpanish() string
}

// Notice that the SpanishSpeaker already implements the
// SpeakSpanish interface
type SpanishSpeaker struct {}

func (s *SpanishSpeaker) greetInSpanish() string {
    return "¡Hola!"
}

type EnglishSpeaker struct {}

func (e *EnglishSpeaker) greetInEnglish() string {
    return "Hello there!"
}
```

As you can see, the `SpanishSpeaker` already implements the `SpeakSpanish` interface. However, the `EnglishSpeaker` does not have a `greetInSpanish` function and hence, doesn't implement the `SpeakSpanish` interface.

Let's use the adapter pattern to write a wrapper over the `EnglishSpeaker` that would give it the ability to greet in spanish.

```go
type EnglishToSpanishAdapter {
    speaker EnglishSpeaker
}

func (a *EnglishToSpanishAdapter) greetInSpanish() string {
    const englishMessage = a.speaker.greetInEnglish();
    return translate(englishMessage)
}

func translate(engMessage string) spanishMessage string {
    // *insert complex translation logic*
    spanishMessage = "¡Hola!"
    return
}
```

Closely observing the adapter will show you that since it is a struct and has the `greetInSpanish` method, it implements the `SpeakSpanish` interface. This is good because once we wrap our english speaker in this adapter, we essentially get a spanish speaker.

_Also notice that we have a translation function as a util along with the adapter. Normally, you would use some service, like google translate, to translate the messages_

Now, our Adapter module is complete. Amazing! 🔥🤘🏽

Let's test it out and see it in action.

```go
func main() {
    espanol := SpanishSpeaker{}
    englishwoman := EnglishSpeaker{}

    fmt.Println("Without translation:")
    fmt.Println("Español says: ", espanol.greetInSpanish())
    fmt.Println("English Woman says: ", englishwoman.greetInEnglish())

    adaptedEnglishwoman := EnglishToSpanishAdapter{
        speaker: englishwoman,
    }

    fmt.Println("------------")
    fmt.Println("With translation:")
    fmt.Println("Español says: ", espanol.greetInSpanish())
    fmt.Println("English Woman says: ", adaptedEnglishwoman.greetInSpanish())
}
```

As a result of this, you should see something like this in the terminal:

```text
Without translation:
Español says:  ¡Hola!
English Woman says:  Hello there
------------
With translation:
Español says:  ¡Hola!
English Woman says:  ¡Hola!
```

Hurray! The translation worked; the english woman and spanish dude can now be friends! 👩🏻‍🤝‍👨🏾

Design patterns save the day, yet again 😁
