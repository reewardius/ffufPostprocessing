package results

import (
	_struct "github.com/Damian89/ffufPostprocessing/pkg/struct"
	"strconv"
)

func MinimizeOriginalResults(Entries *[]_struct.Result) []_struct.Result {

	UniqueStatusMd5 := map[string]int{}
	UniqueStatusSizeMd5 := map[string]int{}
	UniqueStatusWordsMd5 := map[string]int{}
	UniqueStatusLinesMd5 := map[string]int{}
	UniqueStatusContentTypeMd5 := map[string]int{}
	UniqueStatusRedirectAndParameters := map[string]int{}
	UniqueTitleLengthMd5 := map[string]int{}
	UniqueTitleWordsMd5 := map[string]int{}
	UniqueTitleLinesWordsMd5 := map[string]int{}
	UniqueCssFilesMd5 := map[string]int{}
	UniqueJsFilesMd5 := map[string]int{}
	UniqueTagsMd5 := map[string]int{}
	UniqueHttpStatusHeaderCountMd5 := map[string]int{}
	MeanSize := 0

	for i := 0; i < len(*Entries); i++ {
		AnalyzeByHttpStatus(Entries, i, &UniqueStatusMd5)
		AnalyzeByHttpStatusAndSize(Entries, i, &UniqueStatusSizeMd5)
		AnalyzeByHttpStatusAndWords(Entries, i, &UniqueStatusWordsMd5)
		AnalyzeByHttpStatusAndLines(Entries, i, &UniqueStatusLinesMd5)
		AnalyzeByHttpStatusAndContentType(Entries, i, &UniqueStatusContentTypeMd5)
		AnalyzeByHttpStatusAndRedirectData(Entries, i, &UniqueStatusRedirectAndParameters)
		AnalyzeByTitleLength(Entries, i, &UniqueTitleLengthMd5)
		AnalyzeByTitleWords(Entries, i, &UniqueTitleWordsMd5)
		AnalyzeByTitleLengthWords(Entries, i, &UniqueTitleLinesWordsMd5)
		AnalyzeByCssFiles(Entries, i, &UniqueCssFilesMd5)
		AnalyzeByJsFiles(Entries, i, &UniqueJsFilesMd5)
		AnalyzeByTags(Entries, i, &UniqueTagsMd5)
		AnalyzeByHttpStatusAndHeadersCount(Entries, i, &UniqueHttpStatusHeaderCountMd5)
		MeanSize += (*Entries)[i].Length
	}

	MeanSize = MeanSize / len(*Entries)

	TemporaryCleanedResults := []_struct.Result{}
	PositionsDone := map[int]bool{}
	ContentCounterMap := map[string]int{}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "status:" + strconv.Itoa((*Entries)[i].Status)

		if UniqueStatusMd5[Content] > 0 && ContentCounterMap[Content] < 2 {
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "status-size:" + strconv.Itoa((*Entries)[i].Status) + ":" + strconv.Itoa((*Entries)[i].Length)

		if UniqueStatusSizeMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "status-words:" + strconv.Itoa((*Entries)[i].Status) + ":" + strconv.Itoa((*Entries)[i].Words)

		if UniqueStatusWordsMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "status-lines:" + strconv.Itoa((*Entries)[i].Status) + ":" + strconv.Itoa((*Entries)[i].Lines)

		if UniqueStatusLinesMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "status-type:" + strconv.Itoa((*Entries)[i].Status) + ":" + (*Entries)[i].ContentType

		if UniqueStatusContentTypeMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "status-redirect:" + strconv.Itoa((*Entries)[i].Status) + ":" + (*Entries)[i].RedirectDomain + ":" + (*Entries)[i].CountRedirectParameters

		if UniqueStatusRedirectAndParameters[Content] < 5 && ContentCounterMap[Content] < 2 {
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "status-header-count:" + strconv.Itoa((*Entries)[i].Status) + ":" + (*Entries)[i].CountHeaders

		if UniqueHttpStatusHeaderCountMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "title-length:" + (*Entries)[i].LengthTitle

		if UniqueTitleLengthMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "title-words:" + (*Entries)[i].WordsTitle

		if UniqueTitleWordsMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "title-length-words:" + (*Entries)[i].WordsTitle + ":" + (*Entries)[i].LengthTitle

		if UniqueTitleLinesWordsMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "css:" + (*Entries)[i].CountCssFiles

		if UniqueCssFilesMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "js:" + (*Entries)[i].CountJsFiles

		if UniqueJsFilesMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "tags:" + (*Entries)[i].CountTags

		if UniqueTagsMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Dev := (*Entries)[i].Length / MeanSize
		Content := "dev:" + strconv.Itoa(Dev)

		if Dev != 1 && ContentCounterMap[Content] < 2 {
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++

		}

	}

	return TemporaryCleanedResults
}

func AnalyzeByHttpStatus(Entries *[]_struct.Result, i int, StatusMd5 *map[string]int) {

	Content := "status:" + strconv.Itoa((*Entries)[i].Status)

	if (*StatusMd5)[Content] == 0 {
		(*StatusMd5)[Content] = 1
	} else {
		(*StatusMd5)[Content]++
	}

}

func AnalyzeByHttpStatusAndSize(Entries *[]_struct.Result, i int, StatusSizesMd5 *map[string]int) {
	Content := "status-size:" + strconv.Itoa((*Entries)[i].Status) + ":" + strconv.Itoa((*Entries)[i].Length)

	if (*StatusSizesMd5)[Content] == 0 {
		(*StatusSizesMd5)[Content] = 1
	} else {
		(*StatusSizesMd5)[Content]++
	}

}

func AnalyzeByHttpStatusAndHeadersCount(Entries *[]_struct.Result, i int, countMd5 *map[string]int) {
	Content := "status-header-count:" + strconv.Itoa((*Entries)[i].Status) + ":" + (*Entries)[i].CountHeaders

	if (*countMd5)[Content] == 0 {
		(*countMd5)[Content] = 1
	} else {
		(*countMd5)[Content]++
	}

}

func AnalyzeByTags(Entries *[]_struct.Result, i int, tagsMd5 *map[string]int) {
	Content := "tags:" + (*Entries)[i].CountTags

	if (*tagsMd5)[Content] == 0 {
		(*tagsMd5)[Content] = 1
	} else {
		(*tagsMd5)[Content]++
	}

}

func AnalyzeByJsFiles(Entries *[]_struct.Result, i int, filesMd5 *map[string]int) {
	Content := "js:" + (*Entries)[i].CountJsFiles

	if (*filesMd5)[Content] == 0 {
		(*filesMd5)[Content] = 1
	} else {
		(*filesMd5)[Content]++
	}
}

func AnalyzeByCssFiles(Entries *[]_struct.Result, i int, filesMd5 *map[string]int) {
	Content := "css:" + (*Entries)[i].CountCssFiles

	if (*filesMd5)[Content] == 0 {
		(*filesMd5)[Content] = 1
	} else {
		(*filesMd5)[Content]++
	}
}

func AnalyzeByTitleLengthWords(Entries *[]_struct.Result, i int, wordsMd5 *map[string]int) {
	Content := "title-length-words:" + (*Entries)[i].WordsTitle + ":" + (*Entries)[i].LengthTitle

	if (*wordsMd5)[Content] == 0 {
		(*wordsMd5)[Content] = 1
	} else {
		(*wordsMd5)[Content]++
	}
}

func AnalyzeByTitleWords(Entries *[]_struct.Result, i int, wordsMd5 *map[string]int) {
	Content := "title-words:" + (*Entries)[i].WordsTitle

	if (*wordsMd5)[Content] == 0 {
		(*wordsMd5)[Content] = 1
	} else {
		(*wordsMd5)[Content]++
	}
}

func AnalyzeByTitleLength(Entries *[]_struct.Result, i int, lengthMd5 *map[string]int) {
	Content := "title-length:" + (*Entries)[i].LengthTitle

	if (*lengthMd5)[Content] == 0 {
		(*lengthMd5)[Content] = 1
	} else {
		(*lengthMd5)[Content]++
	}
}

func AnalyzeByHttpStatusAndRedirectData(Entries *[]_struct.Result, i int, parameters *map[string]int) {
	Content := "status-redirect:" + strconv.Itoa((*Entries)[i].Status) + ":" + (*Entries)[i].RedirectDomain + ":" + (*Entries)[i].CountRedirectParameters

	if (*parameters)[Content] == 0 {
		(*parameters)[Content] = 1
	} else {
		(*parameters)[Content]++
	}
}

func AnalyzeByHttpStatusAndContentType(Entries *[]_struct.Result, i int, StatusCTMd5 *map[string]int) {
	Content := "status-type:" + strconv.Itoa((*Entries)[i].Status) + ":" + (*Entries)[i].ContentType

	if (*StatusCTMd5)[Content] == 0 {
		(*StatusCTMd5)[Content] = 1
	} else {
		(*StatusCTMd5)[Content]++
	}
}

func AnalyzeByHttpStatusAndLines(Entries *[]_struct.Result, i int, StatusLinesMd5 *map[string]int) {
	Content := "status-lines:" + strconv.Itoa((*Entries)[i].Status) + ":" + strconv.Itoa((*Entries)[i].Lines)

	if (*StatusLinesMd5)[Content] == 0 {
		(*StatusLinesMd5)[Content] = 1
	} else {
		(*StatusLinesMd5)[Content]++
	}

}

func AnalyzeByHttpStatusAndWords(Entries *[]_struct.Result, i int, StatusWordsMd5 *map[string]int) {
	Content := "status-words:" + strconv.Itoa((*Entries)[i].Status) + ":" + strconv.Itoa((*Entries)[i].Words)

	if (*StatusWordsMd5)[Content] == 0 {
		(*StatusWordsMd5)[Content] = 1
	} else {
		(*StatusWordsMd5)[Content]++
	}

}