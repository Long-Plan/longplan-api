package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/Long-Plan/longplan-api/internal/core/domain"
	"github.com/Long-Plan/longplan-api/internal/core/dto"
	"github.com/Long-Plan/longplan-api/pkg/errors"
	"github.com/PuerkitoBio/goquery"
)

type enrolledCourseService struct{}

func NewEnrolledCourseService() domain.EnrolledCourseService {
	return &enrolledCourseService{}
}

func (e *enrolledCourseService) GetEnrolledCoursesByStudentCode(studentCode int) ([]dto.MappingEnrolledCourse, error) {
	body, err := fetchHTTP(studentCode)
	if err != nil {
		return nil, errors.NewNotFoundError("Student ID not found")
	}

	var mappings []dto.MappingEnrolledCourse

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to parse document: %v", err)
	}

	doc.Find("table[width='100%'][border='0'][class='t']").Each(func(i int, s *goquery.Selection) {
		var currentCourses []dto.EnrolledCourse
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
					student_transformed, err := transformInput(studentCode)
					if err != nil {
						return
					}
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

			currentCourses = append(currentCourses, dto.EnrolledCourse{
				CourseNo: courseNo,
				Credit:   credit,
				Grade:    grade,
			})
		})

		if len(currentCourses) > 0 {
			mappings = append(mappings, dto.MappingEnrolledCourse{
				Year:     year,
				Semester: semester,
				Courses:  currentCourses,
			})
		}
	})

	return mappings, nil
}

// Fetch the HTTP response and return as a string
func fetchHTTP(studentCode int) (string, error) {
	url := fmt.Sprintf("https://reg.eng.cmu.ac.th/reg/plan_detail/plan_data_term.php?student_id=%v", studentCode)
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
func transformInput(input int) (string, error) {
	re := regexp.MustCompile(`^(\d{2})`)
	matches := re.FindStringSubmatch(string(input))

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
