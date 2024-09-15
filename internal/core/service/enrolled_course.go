package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/Long-Plan/longplan-api/pkg/errors"
	"github.com/PuerkitoBio/goquery"
)

type EnrolledCourse struct {
	CourseNo string
	Credit   string
	Grade    string
}

type MappingEnrolledCourse struct {
	Year     string
	Semester string
	Courses  []EnrolledCourse
}

type CourseData struct {
	CourseNo       string `json:"courseNo"`
	CourseTitleEng string `json:"CourseTitleEng"`
	Abbreviation   string `json:"Abbreviation"`
}

// Fetch the HTTP response and return as a string
func fetchHTTP(studentID string) (string, error) {
	url := fmt.Sprintf("https://reg.eng.cmu.ac.th/reg/plan_detail/plan_data_term.php?student_id=%v", studentID)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// Extract the first two digits and transform the input string
func transformInput(input string) (string, error) {
	re := regexp.MustCompile(`^(\d{2})`)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 2 {
		return "", fmt.Errorf("Invalid input format")
	}

	firstTwoDigits := matches[1]
	result := "25" + firstTwoDigits

	// Convert to int to remove leading zeros
	resultInt, err := strconv.Atoi(result)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(resultInt), nil
}

// Parse enrolled course data from HTML for a given studentID
func ScrapeEnrollCourse(studentID string) ([]MappingEnrolledCourse, error) {
	body, err := fetchHTTP(studentID)
	if err != nil {
		return nil, errors.NewNotFoundError("Student ID not found")
	}

	var mappings []MappingEnrolledCourse

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to parse document: %v", err)
	}

	doc.Find("table[width='100%'][border='0'][class='t']").Each(func(i int, s *goquery.Selection) {
		var currentCourses []EnrolledCourse
		var year, semester string

		s.Find("td[align='center'] > B").Each(func(j int, row *goquery.Selection) {
			text := row.Text()
			if strings.Contains(text, "ภาคเรียนที่") && strings.Contains(text, "ปีการศึกษา") {
				semesterMatches := regexp.MustCompile(`\d+`).FindAllString(text, -1)
				if len(semesterMatches) >= 2 {
					semester = semesterMatches[len(semesterMatches)-2]
					tempYear := semesterMatches[len(semesterMatches)-1]

					numYear, err := strconv.Atoi(tempYear)
					if err != nil {
						return
					}
					student_transformed, err := transformInput(studentID)
					numStudentID, err := strconv.Atoi(student_transformed)
					if err != nil {
						return
					}

					// Calculate the year based on the student ID
					numYear = numYear - numStudentID + 1
					year = strconv.Itoa(numYear)
				}
			}
		})

		s.Find("table[cellspacing='1'][cellpadding='3'][width='60%'][border='0'][class='t'] tr[bgcolor='#FFFFFF']").Each(func(j int, row *goquery.Selection) {
			courseNo := strings.TrimSpace(row.Find("td:first-child").Text())
			credit := strings.TrimSpace(row.Find("td:nth-child(2)").Text())
			grade := strings.TrimSpace(row.Find("td:nth-child(3)").Text())

			currentCourses = append(currentCourses, EnrolledCourse{
				CourseNo: courseNo,
				Credit:   credit,
				Grade:    grade,
			})
		})

		if len(currentCourses) > 0 {
			mappings = append(mappings, MappingEnrolledCourse{
				Year:     year,
				Semester: semester,
				Courses:  currentCourses,
			})
		}
	})

	return mappings, nil
}

// Fetch the course title using courseID from an API
func fetchCourseTitle(courseID string) (string, error) {
	url := fmt.Sprintf("https://mis-api.cmu.ac.th/tqf/v1/course-template?courseid=%v&academicyear=2563&academicterm=1", courseID)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var courses []CourseData
	err = json.Unmarshal(body, &courses)
	if err != nil {
		return "", err
	}

	if len(courses) == 0 {
		return "", fmt.Errorf("No courses found for courseID %s", courseID)
	}

	if courses[0].Abbreviation != "" {
		return courses[0].Abbreviation, nil
	}
	return courses[0].CourseTitleEng, nil
}

func FilterByGroup(mappings []MappingEnrolledCourse, year string, semester string) []MappingEnrolledCourse {
	var filteredMappings []MappingEnrolledCourse
	for _, mapping := range mappings {
		if mapping.Year == year && mapping.Semester == semester {
			filteredMappings = append(filteredMappings, mapping)
		}
	}
	return filteredMappings
}
