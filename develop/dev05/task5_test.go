package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var testText = `Жил-был в норе под землей хоббит. Не в какой-то там мерзкой
грязной сырой норе, где со всех сторон торчат хвосты червей и
противно пахнет плесенью, но и не в сухой песчаной голой норе, где
не на что сесть и нечего съесть. Нет, нора была хоббичья, а значит —
благоустроенная.
Она начиналась идеально круглой, как иллюминатор, дверью,
выкрашенной зеленой краской, с сияющей медной ручкой точно
посередине. Дверь отворялась внутрь, в длинный коридор, похожий на
железнодорожный туннель, но туннель без гари и без дыма и тоже
очень благоустроенный: стены там были обшиты панелями, пол
выложен плитками и устлан ковром, вдоль стен стояли полированные
стулья, и всюду были прибиты крючочки для шляп и пальто, так как
хоббит любил гостей. Туннель вился все дальше и дальше и заходил
довольно глубоко, но не в самую глубину Холма, как его именовали
жители на много миль в окружности. По обеим сторонам туннеля шли
двери — много-много круглых дверей. Хоббит не признавал
восхождений по лестницам: спальни, ванные, погреба, кладовые (целая
куча кладовых), гардеробные (хоббит отвел несколько комнат под
хранение одежды), кухни, столовые располагались в одном этаже и,
более того, в одном и том же коридоре. Лучшие комнаты находились
по левую руку, и только в них имелись окна — глубоко сидящие
круглые окошечки с видом на сад и на дальние луга, спускавшиеся к
реке.`

func resetFlags() {
	after = 0
	before = 0
	context = 0
	isCount = false
	isIgnore = false
	isInvert = false
	isFixed = false
	isLineNum = false
	grep = ""
}
func TestGrepDefault(t *testing.T) {
	resetFlags()
	var strArrIn []string
	sc := bufio.NewScanner(strings.NewReader(testText))
	for sc.Scan() {
		strArrIn = append(strArrIn, sc.Text())
	}
	result := Grep(strArrIn)
	expect := []string{"Жил-был в норе под землей хоббит. Не в какой-то там мерзкой"}
	require.Equal(t, expect, result)
}

func TestAfterMode(t *testing.T) {
	resetFlags()
	after = 5
	grep = "землей"
	var strArrIn []string
	sc := bufio.NewScanner(strings.NewReader(testText))
	for sc.Scan() {
		strArrIn = append(strArrIn, sc.Text())
	}
	result := Grep(strArrIn)
	expect := []string{"Жил-был в норе под землей хоббит. Не в какой-то там мерзкой",
		"грязной сырой норе, где со всех сторон торчат хвосты червей и",
		"противно пахнет плесенью, но и не в сухой песчаной голой норе, где",
		"не на что сесть и нечего съесть. Нет, нора была хоббичья, а значит —",
		"благоустроенная.",
		"Она начиналась идеально круглой, как иллюминатор, дверью,"}
	require.Equal(t, expect, result)
}

func TestBeforeMode(t *testing.T) {
	resetFlags()
	before = 5
	grep = "сторонам"
	var strArrIn []string
	sc := bufio.NewScanner(strings.NewReader(testText))
	for sc.Scan() {
		strArrIn = append(strArrIn, sc.Text())
	}
	result := Grep(strArrIn)
	expect := []string{"очень благоустроенный: стены там были обшиты панелями, пол",
		"выложен плитками и устлан ковром, вдоль стен стояли полированные",
		"стулья, и всюду были прибиты крючочки для шляп и пальто, так как",
		"хоббит любил гостей. Туннель вился все дальше и дальше и заходил",
		"довольно глубоко, но не в самую глубину Холма, как его именовали",
		"жители на много миль в окружности. По обеим сторонам туннеля шли"}
	require.Equal(t, expect, result)
}

func TestContextMode(t *testing.T) {
	resetFlags()
	context = 1
	grep = "сторонам"
	var strArrIn []string
	sc := bufio.NewScanner(strings.NewReader(testText))
	for sc.Scan() {
		strArrIn = append(strArrIn, sc.Text())
	}
	result := Grep(strArrIn)
	expect := []string{"довольно глубоко, но не в самую глубину Холма, как его именовали",
		"жители на много миль в окружности. По обеим сторонам туннеля шли",
		"двери — много-много круглых дверей. Хоббит не признавал"}
	require.Equal(t, expect, result)
}

func TestCountMode(t *testing.T) {
	resetFlags()
	isCount = true
	grep = "сторонам"
	var strArrIn []string
	sc := bufio.NewScanner(strings.NewReader(testText))
	for sc.Scan() {
		strArrIn = append(strArrIn, sc.Text())
	}
	result := Grep(strArrIn)
	expect := []string{"количество строк: ",
		"23",
		"жители на много миль в окружности. По обеим сторонам туннеля шли"}
	require.Equal(t, expect, result)
}

func TestIgnoreCaseMode(t *testing.T) {
	resetFlags()
	isIgnore = true
	grep = "жил-был"
	var strArrIn []string
	sc := bufio.NewScanner(strings.NewReader(testText))
	for sc.Scan() {
		strArrIn = append(strArrIn, sc.Text())
	}
	result := Grep(strArrIn)
	expect := []string{"Жил-был в норе под землей хоббит. Не в какой-то там мерзкой"}
	require.Equal(t, expect, result)
}

func TestExcludeMode(t *testing.T) {
	resetFlags()
	isInvert = true
	grep = "сторонам"
	var strArrIn []string
	sc := bufio.NewScanner(strings.NewReader(testText))
	for sc.Scan() {
		strArrIn = append(strArrIn, sc.Text())
	}
	result := Grep(strArrIn)
	expect := []string{"Жил-был в норе под землей хоббит. Не в какой-то там мерзкой",
		"грязной сырой норе, где со всех сторон торчат хвосты червей и",
		"противно пахнет плесенью, но и не в сухой песчаной голой норе, где",
		"не на что сесть и нечего съесть. Нет, нора была хоббичья, а значит —",
		"благоустроенная.",
		"Она начиналась идеально круглой, как иллюминатор, дверью,",
		"выкрашенной зеленой краской, с сияющей медной ручкой точно",
		"посередине. Дверь отворялась внутрь, в длинный коридор, похожий на",
		"железнодорожный туннель, но туннель без гари и без дыма и тоже",
		"очень благоустроенный: стены там были обшиты панелями, пол",
		"выложен плитками и устлан ковром, вдоль стен стояли полированные",
		"стулья, и всюду были прибиты крючочки для шляп и пальто, так как",
		"хоббит любил гостей. Туннель вился все дальше и дальше и заходил",
		"довольно глубоко, но не в самую глубину Холма, как его именовали",
		"двери — много-много круглых дверей. Хоббит не признавал",
		"восхождений по лестницам: спальни, ванные, погреба, кладовые (целая",
		"куча кладовых), гардеробные (хоббит отвел несколько комнат под",
		"хранение одежды), кухни, столовые располагались в одном этаже и,",
		"более того, в одном и том же коридоре. Лучшие комнаты находились",
		"по левую руку, и только в них имелись окна — глубоко сидящие",
		"круглые окошечки с видом на сад и на дальние луга, спускавшиеся к",
		"реке."}
	require.Equal(t, expect, result)
}

func TestIsFixedMode(t *testing.T) {
	resetFlags()
	isFixed = true

	t.Run("empty result", func(t *testing.T) {
		grep = "сторонам"
		var strArrIn []string
		sc := bufio.NewScanner(strings.NewReader(testText))
		for sc.Scan() {
			strArrIn = append(strArrIn, sc.Text())
		}
		result := Grep(strArrIn)
		require.Nil(t, result)
	})
	t.Run("full string", func(t *testing.T) {
		grep = "железнодорожный туннель, но туннель без гари и без дыма и тоже"
		var strArrIn []string
		sc := bufio.NewScanner(strings.NewReader(testText))
		for sc.Scan() {
			strArrIn = append(strArrIn, sc.Text())
		}
		result := Grep(strArrIn)
		expect := []string{"железнодорожный туннель, но туннель без гари и без дыма и тоже"}
		require.Equal(t, expect, result)
	})
}

func TestPrintLineNumMode(t *testing.T) {
	resetFlags()
	grep = "железнодорожный"
	isLineNum = true
	var strArrIn []string
	sc := bufio.NewScanner(strings.NewReader(testText))
	for sc.Scan() {
		strArrIn = append(strArrIn, sc.Text())
	}
	result := Grep(strArrIn)
	expect := []string{"номер строки: ", "9", "железнодорожный туннель, но туннель без гари и без дыма и тоже"}
	require.Equal(t, expect, result)
}
