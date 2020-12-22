package Model

type ArticleModel struct {
	Id      int
	Title   string
	Content string
}

func NewArticleModel() *ArticleModel {
	return &ArticleModel{}
}

func (this *ArticleModel) String() string {
	return "ArticleModel"
}
