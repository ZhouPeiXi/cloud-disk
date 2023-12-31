// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UserDetailRequest struct {
	Identity string `json:"identity"`
}

type UserDetailResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type MailCodeSendRequest struct {
	Email string `json:"email"`
}

type MailCodeSendResponse struct {
}

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type UserRegisterResponse struct {
}

type ShareBasicDetailRequest struct {
	Identity string `json:"identity"`
}

type ShareBasicDetailResponse struct {
	Name               string `json:"name"`
	Ext                string `json:"ext"`
	Size               int64  `json:"size"`
	Path               string `json:"path"`
	RepositoryIdentity string `json:"repositoryIdentity"`
}

type FileUploadRequest struct {
	Hash string `json:"hash,optional"`
	Name string `json:"name,optional"`
	Ext  string `json:"ext,optional"`
	Size int64  `json:"size,optional"`
	Path string `json:"path,optional"`
}

type FileUploadResponse struct {
	Identity string `json:"identity"`
	Ext      string `json:"ext"`
	Name     string `json:"name"`
}

type UserRepositorySaveRequest struct {
	ParentId           int    `json:"parentId"`
	RepositoryIdentity string `json:"repositoryIdentity"`
	Ext                string `json:"ext"`
	Name               string `json:"name"`
}

type UserRepositorySaveResponse struct {
}

type UserFilelistRequest struct {
	Id   int64 `json:"id,optional"`
	Page int   `json:"page,optional"`
	Size int   `json:"size,optional"`
}

type UserFilelistResponse struct {
	List  []*UserFile `json:"list"`
	Count int64       `json:"count"`
}

type UserFile struct {
	Id                 int64  `json:"id"`
	Identity           string `json:"identity"`
	RepositoryIdentity string `json:"repositoryIdentity"`
	Name               string `json:"name"`
	Ext                string `json:"ext"`
	Path               string `json:"path"`
	Size               int64  `json:"size"`
}

type UserFileNameUpdateRequest struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
}

type UserFileNameUpdateResponse struct {
}

type UserFolderCreateRequest struct {
	Name     string `json:"name"`
	ParentId int    `json:"parentId"`
}

type UserFolderCreateResponse struct {
	Identity string `json:"identity"`
}

type UserFileDeleteRequest struct {
	Identity string `json:"identity"`
}

type UserFileDeleteResponse struct {
}

type UserFileMoveRequest struct {
	Identity       string `json:"identity"`
	ParentIdentity string `json:"parentIdentity"`
}

type UserFileMoveResponse struct {
}

type ShareBasicCreateRequest struct {
	UserRepositoryIdentity string `json:"userRepositoryIdentity"`
	ExpiredTime            int    `json:"expiredTime"`
}

type ShareBasicCreateResponse struct {
	Identity string `json:"identity"`
}

type ShareBasicSaveRequest struct {
	RepositoryIdentity string `json:"repositoryIdentity"`
	ParentId           int    `json:"parentId"`
}

type ShareBasicSaveResponse struct {
	Identity string `json:"id"`
}

type FileUploadPrepareRequest struct {
	Md5 string `json:"md5"`
	Ext string `json:"ext"`
}

type FileUploadPrepareResponse struct {
	Identity string `json:"identity"`
	UploadId string `json:"uploadId"`
	Key      string `json:"key"`
}

type FileUploadChunkRequest struct {
}

type FileUploadChunkResponse struct {
	ETag string `json:"eTag"` // md5
}

type FileUploadCompleteRequest struct {
	Key        string      `json:"key"`
	UploadId   string      `json:"uploadId"`
	CosObjects []CosObject `json:"cosObjects"`
}

type CosObject struct {
	PartNumber int    `json:"partNumber"`
	ETag       string `json:"eTag"`
}

type FileUploadCompleteResponse struct {
}
