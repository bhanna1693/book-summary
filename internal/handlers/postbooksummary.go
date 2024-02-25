package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bhanna1693/book-summary/internal/templates"
	"github.com/bhanna1693/book-summary/internal/utils"
	"github.com/labstack/echo/v4"
)

type PostBookSummaryHandler struct{}

func NewPostBookSummaryHandler() *PostBookSummaryHandler {
	return &PostBookSummaryHandler{}
}

func (h *PostBookSummaryHandler) ServeHTTP(e echo.Context) error {
	req := new(PostBookSummaryRequest)
	if err := e.Bind(req); err != nil {
		return err
	}

	fmt.Printf("Req: %+v\n", req)

	url := "https://api.openai.com/v1/chat/completions"

	openAiRequest := OpenAIChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []OpenAIChatCompletionRequestMessage{
			{
				Role:    "system",
				Content: fmt.Sprintf("You are a helpful book/novel assistant"),
			},
			{
				Role:    "user",
				Content: fmt.Sprintf("Summarize '%s' by %s. Focus on chapters %s-%s?", req.BookName, req.BookAuthor, req.From, req.To),
			},
		},
	}
	byteArr, err := json.Marshal(openAiRequest)
	if err != nil {
		fmt.Printf("Error marshalling request: %v\n", err)
		return err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(byteArr))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return err
	}

	token := "REMOVE_ME"
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return err
	}

	var openAIResp OpenAIChatCompletionResponse
	if err := json.Unmarshal(body, &openAIResp); err != nil {
		fmt.Printf("Error unmarshalling response JSON: %v\n", err)
		return err
	}

	assistantReply := openAIResp.Choices[0].Message.Content

	fmt.Printf("OpenAI Response: %+v\n", openAIResp)

	return utils.Render(e, templates.BookSummaryDetails(assistantReply))
}

type OpenAIChatCompletionResponse struct {
	Choices []struct {
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
		Message      struct {
			Content string `json:"content"`
			Role    string `json:"role"`
		} `json:"message"`
		Logprobs any `json:"logprobs"`
	} `json:"choices"`
	Created int    `json:"created"`
	ID      string `json:"id"`
	Model   string `json:"model"`
	Object  string `json:"object"`
	Usage   struct {
		CompletionTokens int `json:"completion_tokens"`
		PromptTokens     int `json:"prompt_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type OpenAIChatCompletionRequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIChatCompletionRequest struct {
	Model    string                               `json:"model"`
	Messages []OpenAIChatCompletionRequestMessage `json:"messages"`
}

type PostBookSummaryRequest struct {
	BookName   string `form:"bookName"`
	BookAuthor string `form:"bookAuthor"`
	BookISBN   string `form:"bookISBN"`
	Type       string `form:"type"`
	To         string `form:"to"`
	From       string `form:"from"`
}

type OpenLibraryResp struct {
	NumFound      int  `json:"numFound"`
	Start         int  `json:"start"`
	NumFoundExact bool `json:"numFoundExact"`
	Docs          []struct {
		AuthorAlternativeName []string `json:"author_alternative_name,omitempty"`
		AuthorKey             []string `json:"author_key,omitempty"`
		AuthorName            []string `json:"author_name,omitempty"`
		CoverEditionKey       string   `json:"cover_edition_key,omitempty"`
		CoverI                int      `json:"cover_i,omitempty"`
		EbookAccess           string   `json:"ebook_access"`
		EbookCountI           int      `json:"ebook_count_i"`
		EditionCount          int      `json:"edition_count"`
		EditionKey            []string `json:"edition_key"`
		FirstPublishYear      int      `json:"first_publish_year,omitempty"`
		HasFulltext           bool     `json:"has_fulltext"`
		Isbn                  []string `json:"isbn,omitempty"`
		Key                   string   `json:"key"`
		Language              []string `json:"language,omitempty"`
		LastModifiedI         int      `json:"last_modified_i"`
		Lcc                   []string `json:"lcc,omitempty"`
		NumberOfPagesMedian   int      `json:"number_of_pages_median,omitempty"`
		Oclc                  []string `json:"oclc,omitempty"`
		PublicScanB           bool     `json:"public_scan_b"`
		PublishDate           []string `json:"publish_date,omitempty"`
		PublishPlace          []string `json:"publish_place,omitempty"`
		PublishYear           []int    `json:"publish_year,omitempty"`
		Publisher             []string `json:"publisher,omitempty"`
		Seed                  []string `json:"seed"`
		Title                 string   `json:"title"`
		TitleSort             string   `json:"title_sort"`
		TitleSuggest          string   `json:"title_suggest"`
		Type                  string   `json:"type"`
		Subject               []string `json:"subject,omitempty"`
		RatingsAverage        float64  `json:"ratings_average,omitempty"`
		RatingsSortable       float64  `json:"ratings_sortable,omitempty"`
		RatingsCount          int      `json:"ratings_count,omitempty"`
		RatingsCount1         int      `json:"ratings_count_1,omitempty"`
		RatingsCount2         int      `json:"ratings_count_2,omitempty"`
		RatingsCount3         int      `json:"ratings_count_3,omitempty"`
		RatingsCount4         int      `json:"ratings_count_4,omitempty"`
		RatingsCount5         int      `json:"ratings_count_5,omitempty"`
		ReadinglogCount       int      `json:"readinglog_count,omitempty"`
		WantToReadCount       int      `json:"want_to_read_count,omitempty"`
		CurrentlyReadingCount int      `json:"currently_reading_count,omitempty"`
		AlreadyReadCount      int      `json:"already_read_count,omitempty"`
		PublisherFacet        []string `json:"publisher_facet,omitempty"`
		SubjectFacet          []string `json:"subject_facet,omitempty"`
		Version               int64    `json:"_version_"`
		LccSort               string   `json:"lcc_sort,omitempty"`
		AuthorFacet           []string `json:"author_facet,omitempty"`
		SubjectKey            []string `json:"subject_key,omitempty"`
		Lccn                  []string `json:"lccn,omitempty"`
		Ddc                   []string `json:"ddc,omitempty"`
		Ia                    []string `json:"ia,omitempty"`
		IaCollection          []string `json:"ia_collection,omitempty"`
		IaCollectionS         string   `json:"ia_collection_s,omitempty"`
		LendingEditionS       string   `json:"lending_edition_s,omitempty"`
		LendingIdentifierS    string   `json:"lending_identifier_s,omitempty"`
		PrintdisabledS        string   `json:"printdisabled_s,omitempty"`
		IDGoodreads           []string `json:"id_goodreads,omitempty"`
		IaBoxID               []string `json:"ia_box_id,omitempty"`
		DdcSort               string   `json:"ddc_sort,omitempty"`
		Contributor           []string `json:"contributor,omitempty"`
		IDLibrarything        []string `json:"id_librarything,omitempty"`
		FirstSentence         []string `json:"first_sentence,omitempty"`
		IDOverdrive           []string `json:"id_overdrive,omitempty"`
		IDAmazon              []string `json:"id_amazon,omitempty"`
		IDDnb                 []string `json:"id_dnb,omitempty"`
		Place                 []string `json:"place,omitempty"`
		Time                  []string `json:"time,omitempty"`
		IaLoadedID            []string `json:"ia_loaded_id,omitempty"`
		PlaceKey              []string `json:"place_key,omitempty"`
		TimeFacet             []string `json:"time_facet,omitempty"`
		PlaceFacet            []string `json:"place_facet,omitempty"`
		TimeKey               []string `json:"time_key,omitempty"`
		IDProjectGutenberg    []string `json:"id_project_gutenberg,omitempty"`
		IDLibrivox            []string `json:"id_librivox,omitempty"`
		Person                []string `json:"person,omitempty"`
		IDBetterWorldBooks    []string `json:"id_better_world_books,omitempty"`
		PersonKey             []string `json:"person_key,omitempty"`
		PersonFacet           []string `json:"person_facet,omitempty"`
		Subtitle              string   `json:"subtitle,omitempty"`
	} `json:"docs"`
	NumFound0 int    `json:"num_found"`
	Q         string `json:"q"`
	Offset    any    `json:"offset"`
}
