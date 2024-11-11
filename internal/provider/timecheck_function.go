// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ function.Function = TimeCheckFunction{}
)

func NewTimeCheckFunction() function.Function {
	return TimeCheckFunction{}
}

type TimeCheckFunction struct{}

func (r TimeCheckFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "timecheck"
}

func (r TimeCheckFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether current date and time matches the slot",
		Parameters: []function.Parameter{
			function.ListParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The day list to check",
				Name:               "day",
				ElementType:        types.StringType,
			},
			function.ListParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The month list to check",
				Name:               "month",
				ElementType:        types.StringType,
			},
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The time zone to check",
				Name:               "time_zone",
			},
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The start time to check",
				Name:               "start_time",
			},
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The stop time to check",
				Name:               "stop_time",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r TimeCheckFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var day []string
	var month []string
	var time_zone string
	var start_time string
	var stop_time string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &day, &month, &time_zone, &start_time, &stop_time))
	if resp.Error != nil {
		return
	}

	result, err := timecheck(day, month, time_zone, start_time, stop_time)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}

func timecheck(day []string, month []string, timeZone string, startTime string, stopTime string) (bool, error) {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return false, err
	}

	now := time.Now().In(loc)
	isDayAllowed := false
	for _, d := range day {
		if d == now.Weekday().String() || d == "ALL" {
			isDayAllowed = true
			break
		}
	}
	if !isDayAllowed {
		return false, nil
	}

	isMonthAllowed := false
	for _, d := range month {
		if d == now.Month().String() || d == "ALL" {
			isMonthAllowed = true
			break
		}
	}
	if !isMonthAllowed {
		return false, nil
	}

	start, err := time.ParseInLocation("15:04", startTime, loc)
	if err != nil {
		return false, err
	}

	stop, err := time.ParseInLocation("15:04", stopTime, loc)
	if err != nil {
		return false, err
	}

	currentTime := time.Date(0, 1, 1, now.Hour(), now.Minute(), 0, 0, loc)

	if (currentTime.After(start) || currentTime.Equal(start)) && (currentTime.Before(stop) || currentTime.Equal(stop)) {
		return true, nil
	}
	return false, nil
}
