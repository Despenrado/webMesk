package model

import "encoding/json"

type UserAuth struct {
	ID        string
	SessionId string
	JWTToken  string
}

func (t *UserAuth) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

func (t *UserAuth) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	return nil
}
