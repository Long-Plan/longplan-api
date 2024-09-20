package scraping

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/PuerkitoBio/goquery"
)

func stringPtr(s string) *string {
	if s == "" {
		return nil // Return nil if the string is empty
	}
	return &s
}

func ScrapeCourseDetail(courseNo string) (*model.SysCourseDetail, error) {
	// Fetch the HTML page
	url := fmt.Sprintf("https://mis.cmu.ac.th/tqf/coursepublic.aspx?courseno=%v&semester=2&year=2567", courseNo)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch course details: %v", err)
	}
	defer resp.Body.Close()

	// Parse the HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %v", err)
	}

	// Initialize the course detail object
	var courseDetail model.SysCourseDetail

	// Scrape course number
	courseDetail.CourseNo = strings.TrimSpace(doc.Find("#lblCourseID").Text())

	// Scrape course title (English and Thai)
	courseDetail.TitleLongEN = strings.TrimSpace(doc.Find("#lblCourseTitleEng").Text())
	courseDetail.TitleLongTH = strings.TrimSpace(doc.Find("#lblCourseTitleTha").Text())

	// Scrape course description (English and Thai)
	courseDetail.CourseDescEN = stringPtr(strings.TrimSpace(doc.Find("#lblCourseDescriptionEng").Text()))
	courseDetail.CourseDescTH = stringPtr(strings.TrimSpace(doc.Find("#lblCourseDescriptionTha").Text()))

	// Scrape course credit
	creditText := strings.TrimSpace(doc.Find("#lblCredit").Text())
	creditParts := strings.Split(creditText, "(")
	if len(creditParts) > 0 {
		courseDetail.Credit, _ = strconv.Atoi(strings.TrimSpace(creditParts[0]))
	}

	// Scrape prerequisites
	courseDetail.Prerequisite = stringPtr(strings.TrimSpace(doc.Find("#lblPreequisite").Text()))

	return &courseDetail, nil
}
