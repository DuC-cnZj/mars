package models

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/duc-cnzj/mars/pkg/websocket"

	app "github.com/duc-cnzj/mars/internal/app/helper"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"

	"gorm.io/gorm"
)

type Project struct {
	ID int `json:"id" gorm:"primaryKey;"`

	Name            string `json:"name" gorm:"size:100;not null;comment:'项目名'"`
	GitlabProjectId int    `json:"gitlab_project_id" gorm:"not null;type:integer;"`
	GitlabBranch    string `json:"gitlab_branch" gorm:"not null;size:255;"`
	GitlabCommit    string `json:"gitlab_commit" gorm:"not null;size:255;"`
	Config          string `json:"config"`
	OverrideValues  string `json:"override_values"`
	DockerImage     string `json:"docker_image" gorm:"not null;size:255;default:''"`
	PodSelectors    string `json:"pod_selectors" gorm:"type:text;nullable;"`
	NamespaceId     int    `json:"namespace_id"`
	Atomic          bool   `json:"atomic"`
	ExtraValues     string `json:"extra_values" gorm:"type:text;nullable;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	Namespace Namespace
}

func (project *Project) GetExtraValues() (res []*websocket.ProjectExtraItem) {
	json.Unmarshal([]byte(project.ExtraValues), &res)
	return res
}

func (project *Project) SetPodSelectors(selectors []string) {
	project.PodSelectors = strings.Join(selectors, "|")
}

// GetPodSelectors 不仅包括 deployment 的 pod 还包括其他的 stateful sets...
func (project *Project) GetPodSelectors() []string {
	return strings.Split(project.PodSelectors, "|")
}

func (project *Project) GetAllPods() []v1.Pod {
	var list []v1.Pod
	split := strings.Split(project.PodSelectors, "|")
	if len(split) > 0 {
		for _, labels := range split {
			l, _ := app.K8sClientSet().CoreV1().Pods(project.Namespace.Name).List(context.Background(), metav1.ListOptions{
				LabelSelector: labels,
			})

			list = append(list, l.Items...)
		}
	} else {
		l, _ := app.K8sClientSet().CoreV1().Pods(project.Namespace.Name).List(context.Background(), metav1.ListOptions{
			LabelSelector: "app.kubernetes.io/instance=" + project.Name,
		})
		list = l.Items
	}

	return list
}

func (project *Project) GetAllPodMetrics() []v1beta1.PodMetrics {
	app.DB().Preload("Namespace").First(&project)
	metricses := app.K8sMetrics().MetricsV1beta1().PodMetricses(project.Namespace.Name)
	var list []v1beta1.PodMetrics
	split := strings.Split(project.PodSelectors, "|")
	if len(split) > 0 {
		for _, labels := range split {
			l, _ := metricses.List(context.Background(), metav1.ListOptions{
				LabelSelector: labels,
			})

			list = append(list, l.Items...)
		}
	} else {
		l, _ := metricses.List(context.Background(), metav1.ListOptions{
			LabelSelector: "app.kubernetes.io/instance=" + project.Name,
		})
		list = l.Items
	}

	return list
}
