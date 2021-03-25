//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"
)

// Subscription and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-notifications/2.x#/Subscription
type Subscription struct {
	DBTimestamp    `json:",inline"`
	Id             string    `json:"id,omitempty" validate:"omitempty,uuid"`
	Name           string    `json:"name" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Channels       []Address `json:"channels" validate:"required,gt=0,dive"`
	Receiver       string    `json:"receiver" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Categories     []string  `json:"categories,omitempty" validate:"required_without=Labels,omitempty,gt=0,dive,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Labels         []string  `json:"labels,omitempty" validate:"required_without=Categories,omitempty,gt=0,dive,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Description    string    `json:"description,omitempty"`
	ResendLimit    int64     `json:"resendLimit,omitempty"`
	ResendInterval string    `json:"resendInterval,omitempty" validate:"omitempty,edgex-dto-frequency"`
}

// UpdateSubscription and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-notifications/2.x#/UpdateSubscription
type UpdateSubscription struct {
	Id             *string   `json:"id,omitempty" validate:"required_without=Name,edgex-dto-uuid"`
	Name           *string   `json:"name,omitempty" validate:"required_without=Id,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Channels       []Address `json:"channels,omitempty" validate:"omitempty,dive"`
	Receiver       *string   `json:"receiver,omitempty" validate:"omitempty,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Categories     []string  `json:"categories,omitempty" validate:"omitempty,dive,gt=0,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Labels         []string  `json:"labels,omitempty" validate:"omitempty,dive,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Description    *string   `json:"description,omitempty"`
	ResendLimit    *int64    `json:"resendLimit,omitempty"`
	ResendInterval *string   `json:"resendInterval,omitempty" validate:"omitempty,edgex-dto-frequency"`
}

// ToSubscriptionModel transforms the Subscription DTO to the Subscription Model
func ToSubscriptionModel(s Subscription) models.Subscription {
	var m models.Subscription
	m.Categories = s.Categories
	m.Labels = s.Labels
	m.Channels = ToAddressModels(s.Channels)
	m.DBTimestamp = models.DBTimestamp(s.DBTimestamp)
	m.Description = s.Description
	m.Id = s.Id
	m.Receiver = s.Receiver
	m.Name = s.Name
	m.ResendLimit = s.ResendLimit
	m.ResendInterval = s.ResendInterval
	return m
}

// ToSubscriptionModels transforms the Subscription DTO array to the Subscription model array
func ToSubscriptionModels(subs []Subscription) []models.Subscription {
	models := make([]models.Subscription, len(subs))
	for i, s := range subs {
		models[i] = ToSubscriptionModel(s)
	}
	return models
}

// FromSubscriptionModelToDTO transforms the Subscription Model to the Subscription DTO
func FromSubscriptionModelToDTO(s models.Subscription) Subscription {
	return Subscription{
		DBTimestamp:    DBTimestamp(s.DBTimestamp),
		Categories:     s.Categories,
		Labels:         s.Labels,
		Channels:       FromAddressModelsToDTOs(s.Channels),
		Description:    s.Description,
		Id:             s.Id,
		Receiver:       s.Receiver,
		Name:           s.Name,
		ResendLimit:    s.ResendLimit,
		ResendInterval: s.ResendInterval,
	}
}

// FromSubscriptionModels transforms the Subscription model array to the Subscription DTO array
func FromSubscriptionModelsToDTOs(subscruptions []models.Subscription) []Subscription {
	dtos := make([]Subscription, len(subscruptions))
	for i, s := range subscruptions {
		dtos[i] = FromSubscriptionModelToDTO(s)
	}
	return dtos
}
