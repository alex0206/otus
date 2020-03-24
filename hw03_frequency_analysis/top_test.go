package hw03_frequency_analysis //nolint:golint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Change to true if needed
var taskWithAsteriskIsCompleted = false

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		top, err := Top10("")
		assert.Nil(t, err)
		assert.Len(t, top, 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{"он", "а", "и", "что", "ты", "не", "если", "то", "его", "кристофер", "робин", "в"}
			top, err := Top10(text)
			assert.Nil(t, err)
			assert.Subset(t, expected, top)
		} else {
			expected := []string{"он", "и", "а", "что", "ты", "не", "если", "-", "то", "Кристофер"}
			top, err := Top10(text)
			assert.Nil(t, err)
			assert.ElementsMatch(t, expected, top)
		}
	})

	t.Run("text with less then 10 words", func(t *testing.T) {
		type test struct {
			input    string
			expected []string
		}

		for _, tst := range [...]test{
			{
				input:    "cat and dog, one dog, two cats",
				expected: []string{"cat", "and", "dog,", "one", "two", "cats"},
			},
			{
				input:    "Don't communicate by sharing memory",
				expected: []string{"Don't", "communicate", "by", "sharing", "memory"},
			},
			{
				input:    "word",
				expected: []string{"word"},
			},
		} {
			top, err := Top10(tst.input)
			assert.Nil(t, err)
			assert.ElementsMatch(t, tst.expected, top)
		}
	})
}

func TestWordCntAnalize(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		words, err := wordCntAnalize("")
		assert.Nil(t, err)
		assert.Len(t, words, 0)
	})

	t.Run("analize word count", func(t *testing.T) {
		type test struct {
			input    string
			expected map[string]int
		}

		for _, tst := range [...]test{
			{
				input:    "cat and dog, one dog, two cats",
				expected: map[string]int{"cat": 1, "and": 1, "dog,": 2, "one": 1, "two": 1, "cats": 1},
			},
			{
				input:    "word1 word2 word3 word2 word3 word4 word3 word4 word4 word4",
				expected: map[string]int{"word1": 1, "word2": 2, "word3": 3, "word4": 4},
			},
			{
				input: "Map literals are like struct literals, but the keys are required.",
				expected: map[string]int{
					"Map":       1,
					"literals":  1,
					"are":       2,
					"like":      1,
					"struct":    1,
					"literals,": 1,
					"but":       1,
					"the":       1,
					"keys":      1,
					"required.": 1,
				},
			},
		} {
			words, err := wordCntAnalize(tst.input)
			assert.Nil(t, err)
			assert.True(t, assert.ObjectsAreEqual(tst.expected, words), tst.input)
		}
	})
}

func TestSortWords(t *testing.T) {
	t.Run("empty word cnt map", func(t *testing.T) {
		assert.Len(t, sortWords(nil), 0)
	})

	t.Run("sort words", func(t *testing.T) {
		type test struct {
			input    map[string]int
			expected []word
		}

		for _, tst := range [...]test{
			{
				input: map[string]int{"cat": 10, "and": 1, "dog,": 2, "one": 4, "two": 3},
				expected: []word{
					{"cat", 10},
					{"one", 4},
					{"two", 3},
					{"dog,", 2},
					{"and", 1},
				},
			},
			{
				input: map[string]int{"word1": 31, "word2": 2, "word3": 1, "word4": 4},
				expected: []word{
					{"word1", 31},
					{"word4", 4},
					{"word2", 2},
					{"word3", 1},
				},
			},
			{
				input: map[string]int{
					"Map":       12,
					"literals":  1,
					"are":       2,
					"like":      13,
					"struct":    14,
					"literals,": 11,
					"keys":      15,
					"required.": 16,
				},
				expected: []word{
					{"required.", 16},
					{"keys", 15},
					{"struct", 14},
					{"like", 13},
					{"Map", 12},
					{"literals,", 11},
					{"are", 2},
					{"literals", 1},
				},
			},
		} {
			assert.True(t, assert.ObjectsAreEqual(tst.expected, sortWords(tst.input)), tst.input)
		}
	})
}
