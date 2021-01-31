variable "app_name" {
  description = "Replace Newton codename"
  type        = string
  default     = "newton"
}

variable "app_version" {
  description = "Newton ecosystem semantic version (SemVer)"
  type        = string
  default     = "1.0.0"
}

variable "app_stage" {
  description = "Newton ecosystem stage"
  type        = string
  default     = "prod"
}

variable "webapp_domain" {
  description = "Newton web application domain name"
  type        = string
  default     = "newton.damascus-engineering.com"
}
