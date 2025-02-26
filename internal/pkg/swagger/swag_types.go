package swagger

type SendCodeParams struct {
	Email string `json:"email" example:"some@some.com"`
}

type GetJWTParams struct {
	Email string `json:"email" example:"some@some.com"`
	Code  int    `json:"code" example:"0000"`
}

type GetNewJWTParams struct {
	Refresh_token string `json:"refresh_token" example:"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDMxODYyMzgsImlhdCI6MTc0MDU5NDIzOCwic3ViIjoyLCJ0eXBlIjoicmVmcmVzaCJ9.04yFALX79gM-r9Fj2ZVlUaY0oe3JW-swSzMvt1Sn5bId5rAEiibPghn4QqD3YHWySQnC9PE896j2fvpidEPetP3hZJ9ZhvJdhiNTn-2XVVkH8xyFCQyhB5_BK4dFnDxA2lbkGR14norjHbmokzWBm-ynC132Kh2LN6QmKJpn72t3Xo7ZCTmF4TMCTrfr_ET5zqKKpk343qx1_MefYEvfX9ZTFUC8QvasxfM2wFfbQCkco-qvMC-eAmI3BoIQtXPBnjrY8ifDj54nQi9qgNHX8H7Tzu5tS7AUdkwrkhGU-7oDkY1dYQobBugEL4CAD6TNCwaeduzSTMMB3K_TsgIHHw"`
}
