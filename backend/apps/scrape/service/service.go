package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ariefsn/upwork/env"
	"github.com/ariefsn/upwork/logger"
	"github.com/ariefsn/upwork/models"
	"github.com/ariefsn/upwork/validator"
	"github.com/playwright-community/playwright-go"
)

type scrapeService struct {
}

type scrapeTool struct {
	pw      *playwright.Playwright
	browser playwright.Browser
	page    playwright.Page
}

func (s *scrapeService) tools() (*scrapeTool, error) {
	pw, err := playwright.Run()
	if err != nil {
		logger.Error(err, models.M{
			"file": "scrape_service",
			"func": "tools",
			"line": "pw.Run",
		})
		return nil, err
	}

	browser, err := pw.Firefox.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
		SlowMo:   playwright.Float(0),
		// Timeout:  playwright.Float(5000),
	})
	if err != nil {
		logger.Error(err, models.M{
			"file": "scrape_service",
			"func": "tools",
			"line": "pw.Browser",
		})
		return nil, err
	}

	page, err := browser.NewPage()
	if err != nil {
		logger.Error(err, models.M{
			"file": "scrape_service",
			"func": "tools",
			"line": "pw.Page",
		})
		return nil, err
	}

	return &scrapeTool{
		pw:      pw,
		browser: browser,
		page:    page,
	}, nil
}

// GetProfile implements models.ScrapeService.
func (s *scrapeService) GetProfile(ctx context.Context, userID string) (*models.UpworkProfile, error) {
	err := validator.ValidateVar(userID, "required")
	if err != nil {
		return nil, err
	}

	client, err := s.tools()
	if err != nil {
		return nil, err
	}

	defer func() {
		client.browser.Close()
		if client.pw != nil {
			client.pw.Stop()
		}
	}()

	page := client.page

	url := fmt.Sprintf("%s/%s", env.GetEnv().Urls.UpworkFreelancerProfile, userID)

	if _, err := page.Goto(url, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
		Timeout:   playwright.Float(60000),
	}); err != nil {
		logger.Error(err, models.M{
			"file": "scrape_service",
			"func": "GetProfile",
			"line": "page.Goto",
		})
		return nil, err
	}

	fullName, err := page.Locator(".identity-content [itemprop=name]").
		First().
		InnerText(playwright.LocatorInnerTextOptions{
			// Timeout: playwright.Float(3000),
		})

	if err != nil {
		logger.Error(err, models.M{
			"file": "scrape_service",
			"func": "GetProfile",
			"line": "page.Locator:fullName",
		})
		return nil, err
	}

	fullName = strings.TrimSpace(fullName)

	if fullName == "" {
		logger.Error(err, models.M{
			"file": "scrape_service",
			"func": "GetProfile",
			"line": "fullName.empty",
		})
		return nil, errors.New("invalid user id")
	}

	title, _ := page.Locator(".air3-card-section h2.h4").First().InnerText()
	city, _ := page.Locator(".identity-content [itemprop=locality]").First().InnerText()
	country, _ := page.Locator(".identity-content [itemprop=country-name]").First().InnerText()

	profile := models.UpworkProfile{
		ID:       userID,
		FullName: fullName,
		Title:    title,
		City:     strings.TrimSpace(city),
		Country:  strings.TrimSpace(country),
	}

	return &profile, nil
}

func (s *scrapeService) InstallBrowser() error {
	err := playwright.Install(&playwright.RunOptions{
		Browsers:            []string{"firefox"},
		SkipInstallBrowsers: false,
		Verbose:             true,
	})
	if err != nil {
		logger.Error(err, models.M{
			"file": "scrape_service",
			"func": "tools",
			"line": "pw.Install",
		})
	}
	return err
}

func New() models.ScrapeService {
	return &scrapeService{}
}
