// Code generated by qtc from "basic_search_template.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line basic_search_template.qtpl:1
package server

//line basic_search_template.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line basic_search_template.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line basic_search_template.qtpl:4
type RenderParams struct {
	SearchParams
	NodeId *int64
}

//line basic_search_template.qtpl:10
func StreamRenderScoreSearchBody(qw422016 *qt422016.Writer, params RenderParams) {
//line basic_search_template.qtpl:10
	qw422016.N().S(`
{
    "from": `)
//line basic_search_template.qtpl:12
	qw422016.N().DL(params.From)
//line basic_search_template.qtpl:12
	qw422016.N().S(`,
    "size": `)
//line basic_search_template.qtpl:13
	qw422016.N().DL(params.Size)
//line basic_search_template.qtpl:13
	qw422016.N().S(`,
    "highlight": {
        "order": "score",
        "fragment_size": 80,
        "fields": {
            "title": {
                "number_of_fragments": 1
            },
            "content": {
                "number_of_fragments": 1
            },
            "postscript_list.content": {
                "number_of_fragments": 1
            },
            "reply_list.content": {
                "number_of_fragments": 1,
                "highlight_query": {
                    "nested": {
                        "path": "reply_list",
                        "query": {
                            "match": {
                                "reply_list.content": {
                                    "query": "`)
//line basic_search_template.qtpl:35
	qw422016.N().S(params.Keyword)
//line basic_search_template.qtpl:35
	qw422016.N().S(`",
                                    "analyzer": "ik_smart"
                                }
                            }
                        }
                    }
                }
            }
        }
    },
    "_source": ["title",
                "content",
                "created",
                "id",
                "node",
                "replies",
                "member"],
    "query": {
        "function_score": {
            "query": {
                "bool": {
                    "must": `)
//line basic_search_template.qtpl:56
	streammustQuery(qw422016, params.Gte, params.Lte, params.NodeId)
//line basic_search_template.qtpl:56
	qw422016.N().S(`,
                    "must_not": [
                        {
                            "term": {
                                "deleted": true
                            }
                        }
                    ],
                    "minimum_should_match": 1,
                    "should": [
                        {
                            "match": {
                                "title": {
                                    "query": "`)
//line basic_search_template.qtpl:69
	qw422016.N().S(params.Keyword)
//line basic_search_template.qtpl:69
	qw422016.N().S(`",
                                    "analyzer": "ik_smart",
                                    "boost": 3,
                                    "operator": "`)
//line basic_search_template.qtpl:72
	qw422016.N().S(params.Operator)
//line basic_search_template.qtpl:72
	qw422016.N().S(`"
                                }
                            }
                        },
                        {
                            "bool": {
                                "should": [
                                    {
                                        "match": {
                                            "content": {
                                                "query": "`)
//line basic_search_template.qtpl:82
	qw422016.N().S(params.Keyword)
//line basic_search_template.qtpl:82
	qw422016.N().S(`",
                                                "analyzer": "ik_smart",
                                                "boost": 2,
                                                "operator": "`)
//line basic_search_template.qtpl:85
	qw422016.N().S(params.Operator)
//line basic_search_template.qtpl:85
	qw422016.N().S(`"
                                            }
                                        }
                                    },
                                    {
                                        "nested": {
                                            "path": "postscript_list",
                                            "score_mode": "max",
                                            "query": {
                                                "match": {
                                                    "postscript_list.content": {
                                                        "query": "`)
//line basic_search_template.qtpl:96
	qw422016.N().S(params.Keyword)
//line basic_search_template.qtpl:96
	qw422016.N().S(`",
                                                        "analyzer": "ik_smart",
                                                        "boost": 2,
                                                        "operator": "`)
//line basic_search_template.qtpl:99
	qw422016.N().S(params.Operator)
//line basic_search_template.qtpl:99
	qw422016.N().S(`"
                                                    }
                                                }
                                            }
                                        }
                                    }
                                ]
                            }
                        },
                        {
                            "match": {
                                "all_reply": {
                                    "query": "`)
//line basic_search_template.qtpl:111
	qw422016.N().S(params.Keyword)
//line basic_search_template.qtpl:111
	qw422016.N().S(`",
                                    "analyzer": "ik_smart",
                                    "boost": 1.5,
                                    "operator": "`)
//line basic_search_template.qtpl:114
	qw422016.N().S(params.Operator)
//line basic_search_template.qtpl:114
	qw422016.N().S(`"
                                }
                            }
                        }
                    ]
                }
            },
            "functions": [
                {
                    "filter": {"match_phrase": {
                        "all_content": {
                            "query": "`)
//line basic_search_template.qtpl:125
	qw422016.N().S(params.Keyword)
//line basic_search_template.qtpl:125
	qw422016.N().S(`",
                            "analyzer": "ik_max_word",
                            "slop": 0
                        }
                    }},
                    "weight": 50
                },
                {
                    "field_value_factor": {
                        "field": "bonus",
                        "missing": 0,
                        "modifier": "none",
                        "factor": 1
                    }
                }
            ],
            "score_mode": "sum",
            "boost_mode": "sum"
        }
    }
}
`)
//line basic_search_template.qtpl:146
}

//line basic_search_template.qtpl:146
func WriteRenderScoreSearchBody(qq422016 qtio422016.Writer, params RenderParams) {
//line basic_search_template.qtpl:146
	qw422016 := qt422016.AcquireWriter(qq422016)
//line basic_search_template.qtpl:146
	StreamRenderScoreSearchBody(qw422016, params)
//line basic_search_template.qtpl:146
	qt422016.ReleaseWriter(qw422016)
//line basic_search_template.qtpl:146
}

//line basic_search_template.qtpl:146
func RenderScoreSearchBody(params RenderParams) string {
//line basic_search_template.qtpl:146
	qb422016 := qt422016.AcquireByteBuffer()
//line basic_search_template.qtpl:146
	WriteRenderScoreSearchBody(qb422016, params)
//line basic_search_template.qtpl:146
	qs422016 := string(qb422016.B)
//line basic_search_template.qtpl:146
	qt422016.ReleaseByteBuffer(qb422016)
//line basic_search_template.qtpl:146
	return qs422016
//line basic_search_template.qtpl:146
}

//line basic_search_template.qtpl:149
func StreamRenderTimeOrderSearchBody(qw422016 *qt422016.Writer, params RenderParams) {
//line basic_search_template.qtpl:149
	qw422016.N().S(`
{
    "from": `)
//line basic_search_template.qtpl:151
	qw422016.N().DL(params.From)
//line basic_search_template.qtpl:151
	qw422016.N().S(`,
    "size": `)
//line basic_search_template.qtpl:152
	qw422016.N().DL(params.Size)
//line basic_search_template.qtpl:152
	qw422016.N().S(`,
    "sort": [
        {
            "created": {
                "order":
                `)
//line basic_search_template.qtpl:157
	if params.Order == 1 {
//line basic_search_template.qtpl:157
		qw422016.N().S(`
                "asc"
                `)
//line basic_search_template.qtpl:159
	} else {
//line basic_search_template.qtpl:159
		qw422016.N().S(`
                "desc"
                `)
//line basic_search_template.qtpl:161
	}
//line basic_search_template.qtpl:161
	qw422016.N().S(`
            }
        }
    ],
    "highlight": {
        "order": "score",
        "fragment_size": 80,
        "fields": {
            "title": {
                "number_of_fragments": 1
            },
            "content": {
                "number_of_fragments": 1
            },
            "postscript_list.content": {
                "number_of_fragments": 1
            },
            "reply_list.content": {
                "number_of_fragments": 1,
                "highlight_query": {
                    "nested": {
                        "path": "reply_list",
                        "query": {
                            "match": {
                                "reply_list.content": {
                                    "query": "`)
//line basic_search_template.qtpl:186
	qw422016.N().S(params.Keyword)
//line basic_search_template.qtpl:186
	qw422016.N().S(`",
                                    "analyzer": "ik_smart"
                                }
                            }
                        }
                    }
                }
            }
        }
    },
    "_source": [
        "title",
        "content",
        "created",
        "id",
        "node",
        "replies",
        "member"
    ],
    "query": {
        "constant_score": {
            "filter": {
                "bool": {
                    "must": `)
//line basic_search_template.qtpl:209
	streammustQuery(qw422016, params.Gte, params.Lte, params.NodeId)
//line basic_search_template.qtpl:209
	qw422016.N().S(`,
                    "must_not": [
                        {
                            "term": {
                                "deleted": true
                            }
                        }
                    ],
                    "minimum_should_match": 1,
                    "should": [
                        {
                            "match": {
                                "title": {
                                    "query": "`)
//line basic_search_template.qtpl:222
	qw422016.N().S(params.Keyword)
//line basic_search_template.qtpl:222
	qw422016.N().S(`",
                                    "analyzer": "ik_smart",
                                    "minimum_should_match": "2<70%",
                                    "operator": "`)
//line basic_search_template.qtpl:225
	qw422016.N().S(params.Operator)
//line basic_search_template.qtpl:225
	qw422016.N().S(`"
                                }
                            }
                        },
                        {
                            "match": {
                                "content": {
                                    "query": "`)
//line basic_search_template.qtpl:232
	qw422016.N().S(params.Keyword)
//line basic_search_template.qtpl:232
	qw422016.N().S(`",
                                    "analyzer": "ik_smart",
                                    "minimum_should_match": "2<70%",
                                    "operator": "`)
//line basic_search_template.qtpl:235
	qw422016.N().S(params.Operator)
//line basic_search_template.qtpl:235
	qw422016.N().S(`"
                                }
                            }
                        },
                        {
                            "nested": {
                                "path": "postscript_list",
                                "score_mode": "max",
                                "query": {
                                    "match": {
                                        "postscript_list.content": {
                                            "query": "`)
//line basic_search_template.qtpl:246
	qw422016.N().S(params.Keyword)
//line basic_search_template.qtpl:246
	qw422016.N().S(`",
                                            "analyzer": "ik_smart",
                                            "minimum_should_match": "2<70%",
                                            "operator": "`)
//line basic_search_template.qtpl:249
	qw422016.N().S(params.Operator)
//line basic_search_template.qtpl:249
	qw422016.N().S(`"
                                        }
                                    }
                                }
                            }
                        },
                        {
                            "match_phrase": {
                                "all_reply": {
                                    "query": "`)
//line basic_search_template.qtpl:258
	qw422016.N().S(params.Keyword)
//line basic_search_template.qtpl:258
	qw422016.N().S(`",
                                    "analyzer": "ik_max_word",
                                    "slop": 0
                                }
                            }
                        }
                    ]
                }
            }
        }
    }
}
`)
//line basic_search_template.qtpl:270
}

//line basic_search_template.qtpl:270
func WriteRenderTimeOrderSearchBody(qq422016 qtio422016.Writer, params RenderParams) {
//line basic_search_template.qtpl:270
	qw422016 := qt422016.AcquireWriter(qq422016)
//line basic_search_template.qtpl:270
	StreamRenderTimeOrderSearchBody(qw422016, params)
//line basic_search_template.qtpl:270
	qt422016.ReleaseWriter(qw422016)
//line basic_search_template.qtpl:270
}

//line basic_search_template.qtpl:270
func RenderTimeOrderSearchBody(params RenderParams) string {
//line basic_search_template.qtpl:270
	qb422016 := qt422016.AcquireByteBuffer()
//line basic_search_template.qtpl:270
	WriteRenderTimeOrderSearchBody(qb422016, params)
//line basic_search_template.qtpl:270
	qs422016 := string(qb422016.B)
//line basic_search_template.qtpl:270
	qt422016.ReleaseByteBuffer(qb422016)
//line basic_search_template.qtpl:270
	return qs422016
//line basic_search_template.qtpl:270
}

//line basic_search_template.qtpl:272
func streammustQuery(qw422016 *qt422016.Writer, gte int64, lte int64, nodeId *int64) {
//line basic_search_template.qtpl:272
	qw422016.N().S(`
    `)
//line basic_search_template.qtpl:273
	needComma := false

//line basic_search_template.qtpl:273
	qw422016.N().S(`
    [
        `)
//line basic_search_template.qtpl:275
	if gte > 0 || lte > 0 {
//line basic_search_template.qtpl:275
		qw422016.N().S(`
        `)
//line basic_search_template.qtpl:276
		needComma = true

//line basic_search_template.qtpl:276
		qw422016.N().S(`
        {
          "range": {
            "created": {
              `)
//line basic_search_template.qtpl:280
		if gte > 0 {
//line basic_search_template.qtpl:280
			qw422016.N().S(`"gte": `)
//line basic_search_template.qtpl:280
			qw422016.N().DL(gte)
//line basic_search_template.qtpl:280
			qw422016.N().S(`,`)
//line basic_search_template.qtpl:280
		}
//line basic_search_template.qtpl:280
		qw422016.N().S(`
              `)
//line basic_search_template.qtpl:281
		if lte > 0 {
//line basic_search_template.qtpl:281
			qw422016.N().S(`"lte": `)
//line basic_search_template.qtpl:281
			qw422016.N().DL(lte)
//line basic_search_template.qtpl:281
			qw422016.N().S(`,`)
//line basic_search_template.qtpl:281
		}
//line basic_search_template.qtpl:281
		qw422016.N().S(`
              "format": "epoch_second"
            }
          }
        }
        `)
//line basic_search_template.qtpl:286
	}
//line basic_search_template.qtpl:286
	qw422016.N().S(`
        `)
//line basic_search_template.qtpl:287
	if nodeId != nil {
//line basic_search_template.qtpl:287
		qw422016.N().S(`
        `)
//line basic_search_template.qtpl:288
		if needComma {
//line basic_search_template.qtpl:288
			qw422016.N().S(`,`)
//line basic_search_template.qtpl:288
		}
//line basic_search_template.qtpl:288
		qw422016.N().S(`
        {
            "term": {
                "node": {
                    "value": "`)
//line basic_search_template.qtpl:292
		qw422016.N().DL(*nodeId)
//line basic_search_template.qtpl:292
		qw422016.N().S(`"
                }
            }
        }
        `)
//line basic_search_template.qtpl:296
	}
//line basic_search_template.qtpl:296
	qw422016.N().S(`
    ]
`)
//line basic_search_template.qtpl:298
}

//line basic_search_template.qtpl:298
func writemustQuery(qq422016 qtio422016.Writer, gte int64, lte int64, nodeId *int64) {
//line basic_search_template.qtpl:298
	qw422016 := qt422016.AcquireWriter(qq422016)
//line basic_search_template.qtpl:298
	streammustQuery(qw422016, gte, lte, nodeId)
//line basic_search_template.qtpl:298
	qt422016.ReleaseWriter(qw422016)
//line basic_search_template.qtpl:298
}

//line basic_search_template.qtpl:298
func mustQuery(gte int64, lte int64, nodeId *int64) string {
//line basic_search_template.qtpl:298
	qb422016 := qt422016.AcquireByteBuffer()
//line basic_search_template.qtpl:298
	writemustQuery(qb422016, gte, lte, nodeId)
//line basic_search_template.qtpl:298
	qs422016 := string(qb422016.B)
//line basic_search_template.qtpl:298
	qt422016.ReleaseByteBuffer(qb422016)
//line basic_search_template.qtpl:298
	return qs422016
//line basic_search_template.qtpl:298
}