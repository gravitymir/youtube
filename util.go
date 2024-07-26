package youtube

import "time"

type Format struct {
	Itag             int    `json:"itag"`
	URL              string `json:"url"`
	MimeType         string `json:"mimeType"`
	Bitrate          int    `json:"bitrate"`
	Width            int    `json:"width"`
	Height           int    `json:"height"`
	LastModified     string `json:"lastModified"`
	ContentLength    string `json:"contentLength,omitempty"`
	Quality          string `json:"quality"`
	Fps              int    `json:"fps"`
	QualityLabel     string `json:"qualityLabel"`
	ProjectionType   string `json:"projectionType"`
	AverageBitrate   int    `json:"averageBitrate,omitempty"`
	AudioQuality     string `json:"audioQuality"`
	ApproxDurationMs string `json:"approxDurationMs"`
	AudioSampleRate  string `json:"audioSampleRate"`
	AudioChannels    int    `json:"audioChannels"`
}
type AdaptiveFormat struct {
	Itag      int    `json:"itag"`
	URL       string `json:"url"`
	MimeType  string `json:"mimeType"`
	Bitrate   int    `json:"bitrate"`
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
	InitRange struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"initRange"`
	IndexRange struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"indexRange"`
	LastModified     string  `json:"lastModified"`
	ContentLength    string  `json:"contentLength"`
	Quality          string  `json:"quality"`
	Fps              int     `json:"fps,omitempty"`
	QualityLabel     string  `json:"qualityLabel,omitempty"`
	ProjectionType   string  `json:"projectionType"`
	AverageBitrate   int     `json:"averageBitrate"`
	ApproxDurationMs string  `json:"approxDurationMs"`
	HighReplication  bool    `json:"highReplication,omitempty"`
	AudioQuality     string  `json:"audioQuality,omitempty"`
	AudioSampleRate  string  `json:"audioSampleRate,omitempty"`
	AudioChannels    int     `json:"audioChannels,omitempty"`
	LoudnessDb       float64 `json:"loudnessDb,omitempty"`
}
type StreamingData struct {
	ExpiresInSeconds string           `json:"expiresInSeconds"`
	Formats          []Format         `json:"formats"`
	AdaptiveFormats  []AdaptiveFormat `json:"adaptiveFormats"`
}

type VideoStruct struct {
	ResponseContext struct {
		ServiceTrackingParams []struct {
			Service string `json:"service"`
			Params  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"params"`
		} `json:"serviceTrackingParams"`
		MaxAgeSeconds             int `json:"maxAgeSeconds"`
		MainAppWebResponseContext struct {
			DatasyncID    string `json:"datasyncId"`
			LoggedOut     bool   `json:"loggedOut"`
			TrackingParam string `json:"trackingParam"`
		} `json:"mainAppWebResponseContext"`
		WebResponseContextExtensionData struct {
			HasDecorated bool `json:"hasDecorated"`
		} `json:"webResponseContextExtensionData"`
	} `json:"responseContext"`
	PlayabilityStatus struct {
		Status      string `json:"status"`
		ErrorScreen struct {
			PlayerErrorMessageRenderer struct {
				Reason struct {
					SimpleText string `json:"simpleText"`
				} `json:"reason"`
			} `json:"playerErrorMessageRenderer"`
		} `json:"errorScreen"`
		PlayableInEmbed bool `json:"playableInEmbed"`
		Offlineability  struct {
			OfflineabilityRenderer struct {
				Offlineable bool `json:"offlineable"`
				Formats     []struct {
					Name struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"name"`
					FormatType               string `json:"formatType"`
					AvailabilityType         string `json:"availabilityType"`
					SavedSettingShouldExpire bool   `json:"savedSettingShouldExpire"`
				} `json:"formats"`
				ClickTrackingParams string `json:"clickTrackingParams"`
			} `json:"offlineabilityRenderer"`
		} `json:"offlineability"`
		Miniplayer struct {
			MiniplayerRenderer struct {
				PlaybackMode string `json:"playbackMode"`
			} `json:"miniplayerRenderer"`
		} `json:"miniplayer"`
		ContextParams string `json:"contextParams"`
	} `json:"playabilityStatus"`
	StreamingData   `json:"streamingData"`
	HeartbeatParams struct {
		HeartbeatToken       string `json:"heartbeatToken"`
		IntervalMilliseconds string `json:"intervalMilliseconds"`
		MaxRetries           string `json:"maxRetries"`
		DrmSessionID         string `json:"drmSessionId"`
		SoftFailOnError      bool   `json:"softFailOnError"`
		HeartbeatServerData  string `json:"heartbeatServerData"`
	} `json:"heartbeatParams"`
	PlaybackTracking struct {
		VideostatsPlaybackURL struct {
			BaseURL string `json:"baseUrl"`
		} `json:"videostatsPlaybackUrl"`
		VideostatsDelayplayURL struct {
			BaseURL string `json:"baseUrl"`
		} `json:"videostatsDelayplayUrl"`
		VideostatsWatchtimeURL struct {
			BaseURL string `json:"baseUrl"`
		} `json:"videostatsWatchtimeUrl"`
		PtrackingURL struct {
			BaseURL string `json:"baseUrl"`
		} `json:"ptrackingUrl"`
		QoeURL struct {
			BaseURL string `json:"baseUrl"`
		} `json:"qoeUrl"`
		AtrURL struct {
			BaseURL                 string `json:"baseUrl"`
			ElapsedMediaTimeSeconds int    `json:"elapsedMediaTimeSeconds"`
		} `json:"atrUrl"`
		VideostatsScheduledFlushWalltimeSeconds []int `json:"videostatsScheduledFlushWalltimeSeconds"`
		VideostatsDefaultFlushIntervalSeconds   int   `json:"videostatsDefaultFlushIntervalSeconds"`
	} `json:"playbackTracking"`
	Captions struct {
		PlayerCaptionsTracklistRenderer struct {
			CaptionTracks []CaptionTrack `json:"captionTracks"`
			AudioTracks   []struct {
				CaptionTrackIndices      []int  `json:"captionTrackIndices"`
				DefaultCaptionTrackIndex int    `json:"defaultCaptionTrackIndex"`
				Visibility               string `json:"visibility"`
				HasDefaultTrack          bool   `json:"hasDefaultTrack"`
				CaptionsInitialState     string `json:"captionsInitialState"`
			} `json:"audioTracks"`
			TranslationLanguages []struct {
				LanguageCode string `json:"languageCode"`
				LanguageName struct {
					SimpleText string `json:"simpleText"`
				} `json:"languageName"`
			} `json:"translationLanguages"`
			DefaultAudioTrackIndex int `json:"defaultAudioTrackIndex"`
		} `json:"playerCaptionsTracklistRenderer"`
	} `json:"captions"`
	VideoDetails `json:"videoDetails"`
	Annotations  []struct {
		PlayerAnnotationsExpandedRenderer struct {
			FeaturedChannel struct {
				StartTimeMs string `json:"startTimeMs"`
				EndTimeMs   string `json:"endTimeMs"`
				Watermark   struct {
					Thumbnails []struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"thumbnails"`
				} `json:"watermark"`
				TrackingParams     string `json:"trackingParams"`
				NavigationEndpoint struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							URL         string `json:"url"`
							WebPageType string `json:"webPageType"`
							RootVe      int    `json:"rootVe"`
							APIURL      string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					BrowseEndpoint struct {
						BrowseID string `json:"browseId"`
					} `json:"browseEndpoint"`
				} `json:"navigationEndpoint"`
				ChannelName     string `json:"channelName"`
				SubscribeButton struct {
					SubscribeButtonRenderer struct {
						ButtonText struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"buttonText"`
						Subscribed           bool   `json:"subscribed"`
						Enabled              bool   `json:"enabled"`
						Type                 string `json:"type"`
						ChannelID            string `json:"channelId"`
						ShowPreferences      bool   `json:"showPreferences"`
						SubscribedButtonText struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"subscribedButtonText"`
						UnsubscribedButtonText struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"unsubscribedButtonText"`
						TrackingParams        string `json:"trackingParams"`
						UnsubscribeButtonText struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"unsubscribeButtonText"`
						ServiceEndpoints []struct {
							ClickTrackingParams string `json:"clickTrackingParams"`
							CommandMetadata     struct {
								WebCommandMetadata struct {
									SendPost bool   `json:"sendPost"`
									APIURL   string `json:"apiUrl"`
								} `json:"webCommandMetadata"`
							} `json:"commandMetadata"`
							SubscribeEndpoint struct {
								ChannelIds []string `json:"channelIds"`
								Params     string   `json:"params"`
							} `json:"subscribeEndpoint,omitempty"`
							SignalServiceEndpoint struct {
								Signal  string `json:"signal"`
								Actions []struct {
									ClickTrackingParams string `json:"clickTrackingParams"`
									OpenPopupAction     struct {
										Popup struct {
											ConfirmDialogRenderer struct {
												TrackingParams string `json:"trackingParams"`
												DialogMessages []struct {
													Runs []struct {
														Text string `json:"text"`
													} `json:"runs"`
												} `json:"dialogMessages"`
												ConfirmButton struct {
													ButtonRenderer struct {
														Style      string `json:"style"`
														Size       string `json:"size"`
														IsDisabled bool   `json:"isDisabled"`
														Text       struct {
															Runs []struct {
																Text string `json:"text"`
															} `json:"runs"`
														} `json:"text"`
														ServiceEndpoint struct {
															ClickTrackingParams string `json:"clickTrackingParams"`
															CommandMetadata     struct {
																WebCommandMetadata struct {
																	SendPost bool   `json:"sendPost"`
																	APIURL   string `json:"apiUrl"`
																} `json:"webCommandMetadata"`
															} `json:"commandMetadata"`
															UnsubscribeEndpoint struct {
																ChannelIds []string `json:"channelIds"`
																Params     string   `json:"params"`
															} `json:"unsubscribeEndpoint"`
														} `json:"serviceEndpoint"`
														Accessibility struct {
															Label string `json:"label"`
														} `json:"accessibility"`
														TrackingParams string `json:"trackingParams"`
													} `json:"buttonRenderer"`
												} `json:"confirmButton"`
												CancelButton struct {
													ButtonRenderer struct {
														Style      string `json:"style"`
														Size       string `json:"size"`
														IsDisabled bool   `json:"isDisabled"`
														Text       struct {
															Runs []struct {
																Text string `json:"text"`
															} `json:"runs"`
														} `json:"text"`
														Accessibility struct {
															Label string `json:"label"`
														} `json:"accessibility"`
														TrackingParams string `json:"trackingParams"`
													} `json:"buttonRenderer"`
												} `json:"cancelButton"`
												PrimaryIsCancel bool `json:"primaryIsCancel"`
											} `json:"confirmDialogRenderer"`
										} `json:"popup"`
										PopupType string `json:"popupType"`
									} `json:"openPopupAction"`
								} `json:"actions"`
							} `json:"signalServiceEndpoint,omitempty"`
						} `json:"serviceEndpoints"`
						SubscribeAccessibility struct {
							AccessibilityData struct {
								Label string `json:"label"`
							} `json:"accessibilityData"`
						} `json:"subscribeAccessibility"`
						UnsubscribeAccessibility struct {
							AccessibilityData struct {
								Label string `json:"label"`
							} `json:"accessibilityData"`
						} `json:"unsubscribeAccessibility"`
					} `json:"subscribeButtonRenderer"`
				} `json:"subscribeButton"`
			} `json:"featuredChannel"`
			AllowSwipeDismiss bool   `json:"allowSwipeDismiss"`
			AnnotationID      string `json:"annotationId"`
		} `json:"playerAnnotationsExpandedRenderer"`
	} `json:"annotations"`
	PlayerConfig struct {
		AudioConfig struct {
			LoudnessDb              float64 `json:"loudnessDb"`
			PerceptualLoudnessDb    float64 `json:"perceptualLoudnessDb"`
			EnablePerFormatLoudness bool    `json:"enablePerFormatLoudness"`
		} `json:"audioConfig"`
		StreamSelectionConfig struct {
			MaxBitrate string `json:"maxBitrate"`
		} `json:"streamSelectionConfig"`
		MediaCommonConfig struct {
			DynamicReadaheadConfig struct {
				MaxReadAheadMediaTimeMs int `json:"maxReadAheadMediaTimeMs"`
				MinReadAheadMediaTimeMs int `json:"minReadAheadMediaTimeMs"`
				ReadAheadGrowthRateMs   int `json:"readAheadGrowthRateMs"`
			} `json:"dynamicReadaheadConfig"`
		} `json:"mediaCommonConfig"`
		WebPlayerConfig struct {
			UseCobaltTvosDash       bool `json:"useCobaltTvosDash"`
			WebPlayerActionsPorting struct {
				GetSharePanelCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							SendPost bool   `json:"sendPost"`
							APIURL   string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					WebPlayerShareEntityServiceEndpoint struct {
						SerializedShareEntity string `json:"serializedShareEntity"`
					} `json:"webPlayerShareEntityServiceEndpoint"`
				} `json:"getSharePanelCommand"`
				SubscribeCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							SendPost bool   `json:"sendPost"`
							APIURL   string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					SubscribeEndpoint struct {
						ChannelIds []string `json:"channelIds"`
						Params     string   `json:"params"`
					} `json:"subscribeEndpoint"`
				} `json:"subscribeCommand"`
				UnsubscribeCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							SendPost bool   `json:"sendPost"`
							APIURL   string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					UnsubscribeEndpoint struct {
						ChannelIds []string `json:"channelIds"`
						Params     string   `json:"params"`
					} `json:"unsubscribeEndpoint"`
				} `json:"unsubscribeCommand"`
				AddToWatchLaterCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							SendPost bool   `json:"sendPost"`
							APIURL   string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					PlaylistEditEndpoint struct {
						PlaylistID string `json:"playlistId"`
						Actions    []struct {
							AddedVideoID string `json:"addedVideoId"`
							Action       string `json:"action"`
						} `json:"actions"`
					} `json:"playlistEditEndpoint"`
				} `json:"addToWatchLaterCommand"`
				RemoveFromWatchLaterCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							SendPost bool   `json:"sendPost"`
							APIURL   string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					PlaylistEditEndpoint struct {
						PlaylistID string `json:"playlistId"`
						Actions    []struct {
							Action         string `json:"action"`
							RemovedVideoID string `json:"removedVideoId"`
						} `json:"actions"`
					} `json:"playlistEditEndpoint"`
				} `json:"removeFromWatchLaterCommand"`
			} `json:"webPlayerActionsPorting"`
		} `json:"webPlayerConfig"`
	} `json:"playerConfig"`
	Storyboards struct {
		PlayerStoryboardSpecRenderer struct {
			Spec             string `json:"spec"`
			RecommendedLevel int    `json:"recommendedLevel"`
		} `json:"playerStoryboardSpecRenderer"`
	} `json:"storyboards"`
	Microformat struct {
		PlayerMicroformatRenderer struct {
			Thumbnail struct {
				Thumbnails []struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"thumbnail"`
			Embed struct {
				IframeURL string `json:"iframeUrl"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
			} `json:"embed"`
			Title struct {
				SimpleText string `json:"simpleText"`
			} `json:"title"`
			Description struct {
				SimpleText string `json:"simpleText"`
			} `json:"description"`
			LengthSeconds        string   `json:"lengthSeconds"`
			OwnerProfileURL      string   `json:"ownerProfileUrl"`
			ExternalChannelID    string   `json:"externalChannelId"`
			IsFamilySafe         bool     `json:"isFamilySafe"`
			AvailableCountries   []string `json:"availableCountries"`
			IsUnlisted           bool     `json:"isUnlisted"`
			HasYpcMetadata       bool     `json:"hasYpcMetadata"`
			ViewCount            string   `json:"viewCount"`
			Category             string   `json:"category"`
			PublishDate          string   `json:"publishDate"`
			OwnerChannelName     string   `json:"ownerChannelName"`
			LiveBroadcastDetails struct {
				IsLiveNow      bool      `json:"isLiveNow"`
				StartTimestamp time.Time `json:"startTimestamp"`
				EndTimestamp   time.Time `json:"endTimestamp"`
			} `json:"liveBroadcastDetails"`
			UploadDate string `json:"uploadDate"`
		} `json:"playerMicroformatRenderer"`
	} `json:"microformat"`
	Cards struct {
		CardCollectionRenderer struct {
			Cards []struct {
				CardRenderer struct {
					Teaser struct {
						SimpleCardTeaserRenderer struct {
							Message struct {
								SimpleText string `json:"simpleText"`
							} `json:"message"`
							TrackingParams       string `json:"trackingParams"`
							Prominent            bool   `json:"prominent"`
							LogVisibilityUpdates bool   `json:"logVisibilityUpdates"`
							OnTapCommand         struct {
								ClickTrackingParams                   string `json:"clickTrackingParams"`
								ChangeEngagementPanelVisibilityAction struct {
									TargetID   string `json:"targetId"`
									Visibility string `json:"visibility"`
								} `json:"changeEngagementPanelVisibilityAction"`
							} `json:"onTapCommand"`
						} `json:"simpleCardTeaserRenderer"`
					} `json:"teaser"`
					CueRanges []struct {
						StartCardActiveMs string `json:"startCardActiveMs"`
						EndCardActiveMs   string `json:"endCardActiveMs"`
						TeaserDurationMs  string `json:"teaserDurationMs"`
						IconAfterTeaserMs string `json:"iconAfterTeaserMs"`
					} `json:"cueRanges"`
					TrackingParams string `json:"trackingParams"`
				} `json:"cardRenderer"`
			} `json:"cards"`
			HeaderText struct {
				SimpleText string `json:"simpleText"`
			} `json:"headerText"`
			Icon struct {
				InfoCardIconRenderer struct {
					TrackingParams string `json:"trackingParams"`
				} `json:"infoCardIconRenderer"`
			} `json:"icon"`
			CloseButton struct {
				InfoCardIconRenderer struct {
					TrackingParams string `json:"trackingParams"`
				} `json:"infoCardIconRenderer"`
			} `json:"closeButton"`
			TrackingParams           string `json:"trackingParams"`
			AllowTeaserDismiss       bool   `json:"allowTeaserDismiss"`
			LogIconVisibilityUpdates bool   `json:"logIconVisibilityUpdates"`
		} `json:"cardCollectionRenderer"`
	} `json:"cards"`
	TrackingParams string `json:"trackingParams"`
	Attestation    struct {
		PlayerAttestationRenderer struct {
			Challenge    string `json:"challenge"`
			BotguardData struct {
				Program            string `json:"program"`
				InterpreterSafeURL struct {
					PrivateDoNotAccessOrElseTrustedResourceURLWrappedValue string `json:"privateDoNotAccessOrElseTrustedResourceUrlWrappedValue"`
				} `json:"interpreterSafeUrl"`
				ServerEnvironment int `json:"serverEnvironment"`
			} `json:"botguardData"`
		} `json:"playerAttestationRenderer"`
	} `json:"attestation"`
	AdBreakHeartbeatParams string `json:"adBreakHeartbeatParams"`
	FrameworkUpdates       struct {
		EntityBatchUpdate struct {
			Mutations []struct {
				EntityKey string `json:"entityKey"`
				Type      string `json:"type"`
				Payload   struct {
					OfflineabilityEntity struct {
						Key                     string `json:"key"`
						OfflineabilityRenderer  string `json:"offlineabilityRenderer"`
						AddToOfflineButtonState string `json:"addToOfflineButtonState"`
						ContentCheckOk          bool   `json:"contentCheckOk"`
						RacyCheckOk             bool   `json:"racyCheckOk"`
						LoggingDirectives       struct {
							TrackingParams string `json:"trackingParams"`
							Visibility     struct {
								Types string `json:"types"`
							} `json:"visibility"`
							EnableDisplayloggerExperiment bool `json:"enableDisplayloggerExperiment"`
						} `json:"loggingDirectives"`
					} `json:"offlineabilityEntity"`
				} `json:"payload"`
			} `json:"mutations"`
			Timestamp struct {
				Seconds string `json:"seconds"`
				Nanos   int    `json:"nanos"`
			} `json:"timestamp"`
		} `json:"entityBatchUpdate"`
	} `json:"frameworkUpdates"`
}

type YouTube struct {
	Video struct {
		ID            string
		Status        string // OK ERROR LOGIN_REQUIRED LIVE_STREAM_OFFLINE UNPLAYABLE CONTENT_CHECK_REQUIRED
		Reason        string
		Details       FullDetails
		VideoAndAudio []Format
		VideoFormats  []AdaptiveFormat
		AudioFormats  []AdaptiveFormat
		CaptionTracks []CaptionTrack
		Captions      struct {
			Text []struct {
				Text  string `xml:",chardata"`
				Start string `xml:"start,attr"`
				Dur   string `xml:"dur,attr"`
			} `xml:"text"`
		}
		VideoDataByteSlice []byte
	}
	PlayList struct {
		List
		PlayListDataByteSlice []byte
	}
	CaptionsDataByteSlice []byte
	Captions              struct {
		Text []struct {
			Text  string `xml:",chardata"`
			Start string `xml:"start,attr"`
			Dur   string `xml:"dur,attr"`
		} `xml:"text"`
	}
	VideoAndAudio  []Format
	CaptionsTracks []CaptionTrack
}

//type PlayListStruct struct {
//	ResponseContext struct {
//		ServiceTrackingParams []struct {
//			Service string `json:"service"`
//			Params  []struct {
//				Key   string `json:"key"`
//				Value string `json:"value"`
//			} `json:"params"`
//		} `json:"serviceTrackingParams"`
//		MainAppWebResponseContext struct {
//			LoggedOut     bool   `json:"loggedOut"`
//			TrackingParam string `json:"trackingParam"`
//		} `json:"mainAppWebResponseContext"`
//		WebResponseContextExtensionData struct {
//			YtConfigData struct {
//				VisitorData           string `json:"visitorData"`
//				RootVisualElementType int    `json:"rootVisualElementType"`
//			} `json:"ytConfigData"`
//			HasDecorated bool `json:"hasDecorated"`
//		} `json:"webResponseContextExtensionData"`
//	} `json:"responseContext"`
//	Contents struct {
//		TwoColumnBrowseResultsRenderer struct {
//			Tabs []struct {
//				TabRenderer struct {
//					Selected bool `json:"selected"`
//					Content  struct {
//						SectionListRenderer struct {
//							Contents []struct {
//								ItemSectionRenderer struct {
//									Contents []struct {
//										PlaylistVideoListRenderer struct {
//											Contents []struct {
//												PlaylistVideoRenderer struct {
//													VideoID   string `json:"videoId"`
//													Thumbnail struct {
//														Thumbnails []struct {
//															URL    string `json:"url"`
//															Width  int    `json:"width"`
//															Height int    `json:"height"`
//														} `json:"thumbnails"`
//													} `json:"thumbnail"`
//													Title struct {
//														Runs []struct {
//															Text string `json:"text"`
//														} `json:"runs"`
//														Accessibility struct {
//															AccessibilityData struct {
//																Label string `json:"label"`
//															} `json:"accessibilityData"`
//														} `json:"accessibility"`
//													} `json:"title"`
//													Index struct {
//														SimpleText string `json:"simpleText"`
//													} `json:"index"`
//													ShortBylineText struct {
//														Runs []struct {
//															Text               string `json:"text"`
//															NavigationEndpoint struct {
//																ClickTrackingParams string `json:"clickTrackingParams"`
//																CommandMetadata     struct {
//																	WebCommandMetadata struct {
//																		URL         string `json:"url"`
//																		WebPageType string `json:"webPageType"`
//																		RootVe      int    `json:"rootVe"`
//																		APIURL      string `json:"apiUrl"`
//																	} `json:"webCommandMetadata"`
//																} `json:"commandMetadata"`
//																BrowseEndpoint struct {
//																	BrowseID         string `json:"browseId"`
//																	CanonicalBaseURL string `json:"canonicalBaseUrl"`
//																} `json:"browseEndpoint"`
//															} `json:"navigationEndpoint"`
//														} `json:"runs"`
//													} `json:"shortBylineText"`
//													LengthText struct {
//														Accessibility struct {
//															AccessibilityData struct {
//																Label string `json:"label"`
//															} `json:"accessibilityData"`
//														} `json:"accessibility"`
//														SimpleText string `json:"simpleText"`
//													} `json:"lengthText"`
//													NavigationEndpoint struct {
//														ClickTrackingParams string `json:"clickTrackingParams"`
//														CommandMetadata     struct {
//															WebCommandMetadata struct {
//																URL         string `json:"url"`
//																WebPageType string `json:"webPageType"`
//																RootVe      int    `json:"rootVe"`
//															} `json:"webCommandMetadata"`
//														} `json:"commandMetadata"`
//														WatchEndpoint struct {
//															VideoID        string `json:"videoId"`
//															PlaylistID     string `json:"playlistId"`
//															Index          int    `json:"index"`
//															Params         string `json:"params"`
//															PlayerParams   string `json:"playerParams"`
//															LoggingContext struct {
//																VssLoggingContext struct {
//																	SerializedContextData string `json:"serializedContextData"`
//																} `json:"vssLoggingContext"`
//															} `json:"loggingContext"`
//															WatchEndpointSupportedOnesieConfig struct {
//																HTML5PlaybackOnesieConfig struct {
//																	CommonConfig struct {
//																		URL string `json:"url"`
//																	} `json:"commonConfig"`
//																} `json:"html5PlaybackOnesieConfig"`
//															} `json:"watchEndpointSupportedOnesieConfig"`
//														} `json:"watchEndpoint"`
//													} `json:"navigationEndpoint"`
//													LengthSeconds  string `json:"lengthSeconds"`
//													TrackingParams string `json:"trackingParams"`
//													IsPlayable     bool   `json:"isPlayable"`
//													Menu           struct {
//														MenuRenderer struct {
//															Items []struct {
//																MenuServiceItemRenderer struct {
//																	Text struct {
//																		Runs []struct {
//																			Text string `json:"text"`
//																		} `json:"runs"`
//																	} `json:"text"`
//																	Icon struct {
//																		IconType string `json:"iconType"`
//																	} `json:"icon"`
//																	ServiceEndpoint struct {
//																		ClickTrackingParams string `json:"clickTrackingParams"`
//																		CommandMetadata     struct {
//																			WebCommandMetadata struct {
//																				SendPost bool `json:"sendPost"`
//																			} `json:"webCommandMetadata"`
//																		} `json:"commandMetadata"`
//																		SignalServiceEndpoint struct {
//																			Signal  string `json:"signal"`
//																			Actions []struct {
//																				ClickTrackingParams  string `json:"clickTrackingParams"`
//																				AddToPlaylistCommand struct {
//																					OpenMiniplayer      bool   `json:"openMiniplayer"`
//																					VideoID             string `json:"videoId"`
//																					ListType            string `json:"listType"`
//																					OnCreateListCommand struct {
//																						ClickTrackingParams string `json:"clickTrackingParams"`
//																						CommandMetadata     struct {
//																							WebCommandMetadata struct {
//																								SendPost bool   `json:"sendPost"`
//																								APIURL   string `json:"apiUrl"`
//																							} `json:"webCommandMetadata"`
//																						} `json:"commandMetadata"`
//																						CreatePlaylistServiceEndpoint struct {
//																							VideoIds []string `json:"videoIds"`
//																							Params   string   `json:"params"`
//																						} `json:"createPlaylistServiceEndpoint"`
//																					} `json:"onCreateListCommand"`
//																					VideoIds []string `json:"videoIds"`
//																				} `json:"addToPlaylistCommand"`
//																			} `json:"actions"`
//																		} `json:"signalServiceEndpoint"`
//																	} `json:"serviceEndpoint"`
//																	TrackingParams string `json:"trackingParams"`
//																} `json:"menuServiceItemRenderer,omitempty"`
//																MenuServiceItemRenderer0 struct {
//																	Text struct {
//																		Runs []struct {
//																			Text string `json:"text"`
//																		} `json:"runs"`
//																	} `json:"text"`
//																	Icon struct {
//																		IconType string `json:"iconType"`
//																	} `json:"icon"`
//																	ServiceEndpoint struct {
//																		ClickTrackingParams string `json:"clickTrackingParams"`
//																		CommandMetadata     struct {
//																			WebCommandMetadata struct {
//																				SendPost bool   `json:"sendPost"`
//																				APIURL   string `json:"apiUrl"`
//																			} `json:"webCommandMetadata"`
//																		} `json:"commandMetadata"`
//																		ShareEntityServiceEndpoint struct {
//																			SerializedShareEntity string `json:"serializedShareEntity"`
//																			Commands              []struct {
//																				ClickTrackingParams string `json:"clickTrackingParams"`
//																				OpenPopupAction     struct {
//																					Popup struct {
//																						UnifiedSharePanelRenderer struct {
//																							TrackingParams     string `json:"trackingParams"`
//																							ShowLoadingSpinner bool   `json:"showLoadingSpinner"`
//																						} `json:"unifiedSharePanelRenderer"`
//																					} `json:"popup"`
//																					PopupType string `json:"popupType"`
//																					BeReused  bool   `json:"beReused"`
//																				} `json:"openPopupAction"`
//																			} `json:"commands"`
//																		} `json:"shareEntityServiceEndpoint"`
//																	} `json:"serviceEndpoint"`
//																	TrackingParams string `json:"trackingParams"`
//																	HasSeparator   bool   `json:"hasSeparator"`
//																} `json:"menuServiceItemRenderer,omitempty"`
//															} `json:"items"`
//															TrackingParams string `json:"trackingParams"`
//															Accessibility  struct {
//																AccessibilityData struct {
//																	Label string `json:"label"`
//																} `json:"accessibilityData"`
//															} `json:"accessibility"`
//														} `json:"menuRenderer"`
//													} `json:"menu"`
//													ThumbnailOverlays []struct {
//														ThumbnailOverlayTimeStatusRenderer struct {
//															Text struct {
//																Accessibility struct {
//																	AccessibilityData struct {
//																		Label string `json:"label"`
//																	} `json:"accessibilityData"`
//																} `json:"accessibility"`
//																SimpleText string `json:"simpleText"`
//															} `json:"text"`
//															Style string `json:"style"`
//														} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
//														ThumbnailOverlayNowPlayingRenderer struct {
//															Text struct {
//																Runs []struct {
//																	Text string `json:"text"`
//																} `json:"runs"`
//															} `json:"text"`
//														} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
//													} `json:"thumbnailOverlays"`
//													VideoInfo struct {
//														Runs []struct {
//															Text string `json:"text"`
//														} `json:"runs"`
//													} `json:"videoInfo"`
//												} `json:"playlistVideoRenderer,omitempty"`
//												ContinuationItemRenderer struct {
//													Trigger              string `json:"trigger"`
//													ContinuationEndpoint struct {
//														ClickTrackingParams string `json:"clickTrackingParams"`
//														CommandMetadata     struct {
//															WebCommandMetadata struct {
//																SendPost bool   `json:"sendPost"`
//																APIURL   string `json:"apiUrl"`
//															} `json:"webCommandMetadata"`
//														} `json:"commandMetadata"`
//														ContinuationCommand struct {
//															Token   string `json:"token"`
//															Request string `json:"request"`
//														} `json:"continuationCommand"`
//													} `json:"continuationEndpoint"`
//												} `json:"continuationItemRenderer,omitempty"`
//											} `json:"contents"`
//											PlaylistID     string `json:"playlistId"`
//											IsEditable     bool   `json:"isEditable"`
//											CanReorder     bool   `json:"canReorder"`
//											TrackingParams string `json:"trackingParams"`
//											TargetID       string `json:"targetId"`
//										} `json:"playlistVideoListRenderer"`
//									} `json:"contents"`
//									TrackingParams string `json:"trackingParams"`
//								} `json:"itemSectionRenderer,omitempty"`
//								ContinuationItemRenderer struct {
//									Trigger              string `json:"trigger"`
//									ContinuationEndpoint struct {
//										ClickTrackingParams string `json:"clickTrackingParams"`
//										CommandMetadata     struct {
//											WebCommandMetadata struct {
//												SendPost bool   `json:"sendPost"`
//												APIURL   string `json:"apiUrl"`
//											} `json:"webCommandMetadata"`
//										} `json:"commandMetadata"`
//										ContinuationCommand struct {
//											Token   string `json:"token"`
//											Request string `json:"request"`
//										} `json:"continuationCommand"`
//									} `json:"continuationEndpoint"`
//								} `json:"continuationItemRenderer,omitempty"`
//							} `json:"contents"`
//							TrackingParams string `json:"trackingParams"`
//							TargetID       string `json:"targetId"`
//						} `json:"sectionListRenderer"`
//					} `json:"content"`
//					TrackingParams string `json:"trackingParams"`
//				} `json:"tabRenderer"`
//			} `json:"tabs"`
//		} `json:"twoColumnBrowseResultsRenderer"`
//	} `json:"contents"`
//	Header struct {
//		PlaylistHeaderRenderer struct {
//			PlaylistID string `json:"playlistId"`
//			Title      struct {
//				SimpleText string `json:"simpleText"`
//			} `json:"title"`
//			NumVideosText struct {
//				Runs []struct {
//					Text string `json:"text"`
//				} `json:"runs"`
//			} `json:"numVideosText"`
//			DescriptionText struct {
//				SimpleText string `json:"simpleText"`
//			} `json:"descriptionText"`
//			OwnerText struct {
//				Runs []struct {
//					Text               string `json:"text"`
//					NavigationEndpoint struct {
//						ClickTrackingParams string `json:"clickTrackingParams"`
//						CommandMetadata     struct {
//							WebCommandMetadata struct {
//								URL         string `json:"url"`
//								WebPageType string `json:"webPageType"`
//								RootVe      int    `json:"rootVe"`
//								APIURL      string `json:"apiUrl"`
//							} `json:"webCommandMetadata"`
//						} `json:"commandMetadata"`
//						BrowseEndpoint struct {
//							BrowseID         string `json:"browseId"`
//							CanonicalBaseURL string `json:"canonicalBaseUrl"`
//						} `json:"browseEndpoint"`
//					} `json:"navigationEndpoint"`
//				} `json:"runs"`
//			} `json:"ownerText"`
//			ViewCountText struct {
//				SimpleText string `json:"simpleText"`
//			} `json:"viewCountText"`
//			ShareData struct {
//				CanShare bool `json:"canShare"`
//			} `json:"shareData"`
//			IsEditable    bool   `json:"isEditable"`
//			Privacy       string `json:"privacy"`
//			OwnerEndpoint struct {
//				ClickTrackingParams string `json:"clickTrackingParams"`
//				CommandMetadata     struct {
//					WebCommandMetadata struct {
//						URL         string `json:"url"`
//						WebPageType string `json:"webPageType"`
//						RootVe      int    `json:"rootVe"`
//						APIURL      string `json:"apiUrl"`
//					} `json:"webCommandMetadata"`
//				} `json:"commandMetadata"`
//				BrowseEndpoint struct {
//					BrowseID         string `json:"browseId"`
//					CanonicalBaseURL string `json:"canonicalBaseUrl"`
//				} `json:"browseEndpoint"`
//			} `json:"ownerEndpoint"`
//			EditableDetails struct {
//				CanDelete bool `json:"canDelete"`
//			} `json:"editableDetails"`
//			TrackingParams   string `json:"trackingParams"`
//			ServiceEndpoints []struct {
//				ClickTrackingParams string `json:"clickTrackingParams"`
//				CommandMetadata     struct {
//					WebCommandMetadata struct {
//						SendPost bool   `json:"sendPost"`
//						APIURL   string `json:"apiUrl"`
//					} `json:"webCommandMetadata"`
//				} `json:"commandMetadata"`
//				PlaylistEditEndpoint struct {
//					Actions []struct {
//						Action           string `json:"action"`
//						SourcePlaylistID string `json:"sourcePlaylistId"`
//					} `json:"actions"`
//				} `json:"playlistEditEndpoint"`
//			} `json:"serviceEndpoints"`
//			Stats []struct {
//				Runs []struct {
//					Text string `json:"text"`
//				} `json:"runs,omitempty"`
//				SimpleText string `json:"simpleText,omitempty"`
//			} `json:"stats"`
//			BriefStats []struct {
//				Runs []struct {
//					Text string `json:"text"`
//				} `json:"runs"`
//			} `json:"briefStats"`
//			PlaylistHeaderBanner struct {
//				HeroPlaylistThumbnailRenderer struct {
//					Thumbnail struct {
//						Thumbnails []struct {
//							URL    string `json:"url"`
//							Width  int    `json:"width"`
//							Height int    `json:"height"`
//						} `json:"thumbnails"`
//					} `json:"thumbnail"`
//					MaxRatio       float64 `json:"maxRatio"`
//					TrackingParams string  `json:"trackingParams"`
//					OnTap          struct {
//						ClickTrackingParams string `json:"clickTrackingParams"`
//						CommandMetadata     struct {
//							WebCommandMetadata struct {
//								URL         string `json:"url"`
//								WebPageType string `json:"webPageType"`
//								RootVe      int    `json:"rootVe"`
//							} `json:"webCommandMetadata"`
//						} `json:"commandMetadata"`
//						WatchEndpoint struct {
//							VideoID        string `json:"videoId"`
//							PlaylistID     string `json:"playlistId"`
//							PlayerParams   string `json:"playerParams"`
//							LoggingContext struct {
//								VssLoggingContext struct {
//									SerializedContextData string `json:"serializedContextData"`
//								} `json:"vssLoggingContext"`
//							} `json:"loggingContext"`
//							WatchEndpointSupportedOnesieConfig struct {
//								HTML5PlaybackOnesieConfig struct {
//									CommonConfig struct {
//										URL string `json:"url"`
//									} `json:"commonConfig"`
//								} `json:"html5PlaybackOnesieConfig"`
//							} `json:"watchEndpointSupportedOnesieConfig"`
//						} `json:"watchEndpoint"`
//					} `json:"onTap"`
//					ThumbnailOverlays struct {
//						ThumbnailOverlayHoverTextRenderer struct {
//							Text struct {
//								SimpleText string `json:"simpleText"`
//							} `json:"text"`
//							Icon struct {
//								IconType string `json:"iconType"`
//							} `json:"icon"`
//						} `json:"thumbnailOverlayHoverTextRenderer"`
//					} `json:"thumbnailOverlays"`
//				} `json:"heroPlaylistThumbnailRenderer"`
//			} `json:"playlistHeaderBanner"`
//			SaveButton struct {
//				ToggleButtonRenderer struct {
//					Style struct {
//						StyleType string `json:"styleType"`
//					} `json:"style"`
//					Size struct {
//						SizeType string `json:"sizeType"`
//					} `json:"size"`
//					IsToggled   bool `json:"isToggled"`
//					IsDisabled  bool `json:"isDisabled"`
//					DefaultIcon struct {
//						IconType string `json:"iconType"`
//					} `json:"defaultIcon"`
//					ToggledIcon struct {
//						IconType string `json:"iconType"`
//					} `json:"toggledIcon"`
//					TrackingParams string `json:"trackingParams"`
//					DefaultTooltip string `json:"defaultTooltip"`
//					ToggledTooltip string `json:"toggledTooltip"`
//					ToggledStyle   struct {
//						StyleType string `json:"styleType"`
//					} `json:"toggledStyle"`
//					DefaultNavigationEndpoint struct {
//						ClickTrackingParams string `json:"clickTrackingParams"`
//						CommandMetadata     struct {
//							WebCommandMetadata struct {
//								IgnoreNavigation bool `json:"ignoreNavigation"`
//							} `json:"webCommandMetadata"`
//						} `json:"commandMetadata"`
//						ModalEndpoint struct {
//							Modal struct {
//								ModalWithTitleAndButtonRenderer struct {
//									Title struct {
//										SimpleText string `json:"simpleText"`
//									} `json:"title"`
//									Content struct {
//										SimpleText string `json:"simpleText"`
//									} `json:"content"`
//									Button struct {
//										ButtonRenderer struct {
//											Style      string `json:"style"`
//											Size       string `json:"size"`
//											IsDisabled bool   `json:"isDisabled"`
//											Text       struct {
//												SimpleText string `json:"simpleText"`
//											} `json:"text"`
//											NavigationEndpoint struct {
//												ClickTrackingParams string `json:"clickTrackingParams"`
//												CommandMetadata     struct {
//													WebCommandMetadata struct {
//														URL         string `json:"url"`
//														WebPageType string `json:"webPageType"`
//														RootVe      int    `json:"rootVe"`
//													} `json:"webCommandMetadata"`
//												} `json:"commandMetadata"`
//												SignInEndpoint struct {
//													NextEndpoint struct {
//														ClickTrackingParams string `json:"clickTrackingParams"`
//														CommandMetadata     struct {
//															WebCommandMetadata struct {
//																URL         string `json:"url"`
//																WebPageType string `json:"webPageType"`
//																RootVe      int    `json:"rootVe"`
//																APIURL      string `json:"apiUrl"`
//															} `json:"webCommandMetadata"`
//														} `json:"commandMetadata"`
//														BrowseEndpoint struct {
//															BrowseID string `json:"browseId"`
//														} `json:"browseEndpoint"`
//													} `json:"nextEndpoint"`
//													IdamTag string `json:"idamTag"`
//												} `json:"signInEndpoint"`
//											} `json:"navigationEndpoint"`
//											TrackingParams string `json:"trackingParams"`
//										} `json:"buttonRenderer"`
//									} `json:"button"`
//								} `json:"modalWithTitleAndButtonRenderer"`
//							} `json:"modal"`
//						} `json:"modalEndpoint"`
//					} `json:"defaultNavigationEndpoint"`
//					AccessibilityData struct {
//						AccessibilityData struct {
//							Label string `json:"label"`
//						} `json:"accessibilityData"`
//					} `json:"accessibilityData"`
//					ToggledAccessibilityData struct {
//						AccessibilityData struct {
//							Label string `json:"label"`
//						} `json:"accessibilityData"`
//					} `json:"toggledAccessibilityData"`
//				} `json:"toggleButtonRenderer"`
//			} `json:"saveButton"`
//			ShareButton struct {
//				ButtonRenderer struct {
//					Style           string `json:"style"`
//					Size            string `json:"size"`
//					IsDisabled      bool   `json:"isDisabled"`
//					ServiceEndpoint struct {
//						ClickTrackingParams string `json:"clickTrackingParams"`
//						CommandMetadata     struct {
//							WebCommandMetadata struct {
//								SendPost bool   `json:"sendPost"`
//								APIURL   string `json:"apiUrl"`
//							} `json:"webCommandMetadata"`
//						} `json:"commandMetadata"`
//						ShareEntityServiceEndpoint struct {
//							SerializedShareEntity string `json:"serializedShareEntity"`
//							Commands              []struct {
//								ClickTrackingParams string `json:"clickTrackingParams"`
//								OpenPopupAction     struct {
//									Popup struct {
//										UnifiedSharePanelRenderer struct {
//											TrackingParams     string `json:"trackingParams"`
//											ShowLoadingSpinner bool   `json:"showLoadingSpinner"`
//										} `json:"unifiedSharePanelRenderer"`
//									} `json:"popup"`
//									PopupType string `json:"popupType"`
//									BeReused  bool   `json:"beReused"`
//								} `json:"openPopupAction"`
//							} `json:"commands"`
//						} `json:"shareEntityServiceEndpoint"`
//					} `json:"serviceEndpoint"`
//					Icon struct {
//						IconType string `json:"iconType"`
//					} `json:"icon"`
//					Tooltip           string `json:"tooltip"`
//					TrackingParams    string `json:"trackingParams"`
//					AccessibilityData struct {
//						AccessibilityData struct {
//							Label string `json:"label"`
//						} `json:"accessibilityData"`
//					} `json:"accessibilityData"`
//				} `json:"buttonRenderer"`
//			} `json:"shareButton"`
//			MoreActionsMenu struct {
//				MenuRenderer struct {
//					Items []struct {
//						MenuNavigationItemRenderer struct {
//							Text struct {
//								SimpleText string `json:"simpleText"`
//							} `json:"text"`
//							Icon struct {
//								IconType string `json:"iconType"`
//							} `json:"icon"`
//							NavigationEndpoint struct {
//								ClickTrackingParams string `json:"clickTrackingParams"`
//								CommandMetadata     struct {
//									WebCommandMetadata struct {
//										URL         string `json:"url"`
//										WebPageType string `json:"webPageType"`
//										RootVe      int    `json:"rootVe"`
//										APIURL      string `json:"apiUrl"`
//									} `json:"webCommandMetadata"`
//								} `json:"commandMetadata"`
//								BrowseEndpoint struct {
//									BrowseID       string `json:"browseId"`
//									Params         string `json:"params"`
//									Nofollow       bool   `json:"nofollow"`
//									NavigationType string `json:"navigationType"`
//								} `json:"browseEndpoint"`
//							} `json:"navigationEndpoint"`
//							TrackingParams string `json:"trackingParams"`
//						} `json:"menuNavigationItemRenderer"`
//					} `json:"items"`
//					TrackingParams string `json:"trackingParams"`
//					Accessibility  struct {
//						AccessibilityData struct {
//							Label string `json:"label"`
//						} `json:"accessibilityData"`
//					} `json:"accessibility"`
//					TargetID string `json:"targetId"`
//				} `json:"menuRenderer"`
//			} `json:"moreActionsMenu"`
//			PlayButton struct {
//				ButtonRenderer struct {
//					Style      string `json:"style"`
//					Size       string `json:"size"`
//					IsDisabled bool   `json:"isDisabled"`
//					Text       struct {
//						SimpleText string `json:"simpleText"`
//					} `json:"text"`
//					Icon struct {
//						IconType string `json:"iconType"`
//					} `json:"icon"`
//					NavigationEndpoint struct {
//						ClickTrackingParams string `json:"clickTrackingParams"`
//						CommandMetadata     struct {
//							WebCommandMetadata struct {
//								URL         string `json:"url"`
//								WebPageType string `json:"webPageType"`
//								RootVe      int    `json:"rootVe"`
//							} `json:"webCommandMetadata"`
//						} `json:"commandMetadata"`
//						WatchEndpoint struct {
//							VideoID        string `json:"videoId"`
//							PlaylistID     string `json:"playlistId"`
//							PlayerParams   string `json:"playerParams"`
//							LoggingContext struct {
//								VssLoggingContext struct {
//									SerializedContextData string `json:"serializedContextData"`
//								} `json:"vssLoggingContext"`
//							} `json:"loggingContext"`
//							WatchEndpointSupportedOnesieConfig struct {
//								HTML5PlaybackOnesieConfig struct {
//									CommonConfig struct {
//										URL string `json:"url"`
//									} `json:"commonConfig"`
//								} `json:"html5PlaybackOnesieConfig"`
//							} `json:"watchEndpointSupportedOnesieConfig"`
//						} `json:"watchEndpoint"`
//					} `json:"navigationEndpoint"`
//					TrackingParams string `json:"trackingParams"`
//				} `json:"buttonRenderer"`
//			} `json:"playButton"`
//			ShufflePlayButton struct {
//				ButtonRenderer struct {
//					Style      string `json:"style"`
//					Size       string `json:"size"`
//					IsDisabled bool   `json:"isDisabled"`
//					Text       struct {
//						SimpleText string `json:"simpleText"`
//					} `json:"text"`
//					Icon struct {
//						IconType string `json:"iconType"`
//					} `json:"icon"`
//					NavigationEndpoint struct {
//						ClickTrackingParams string `json:"clickTrackingParams"`
//						CommandMetadata     struct {
//							WebCommandMetadata struct {
//								URL         string `json:"url"`
//								WebPageType string `json:"webPageType"`
//								RootVe      int    `json:"rootVe"`
//							} `json:"webCommandMetadata"`
//						} `json:"commandMetadata"`
//						WatchEndpoint struct {
//							VideoID        string `json:"videoId"`
//							PlaylistID     string `json:"playlistId"`
//							Params         string `json:"params"`
//							PlayerParams   string `json:"playerParams"`
//							LoggingContext struct {
//								VssLoggingContext struct {
//									SerializedContextData string `json:"serializedContextData"`
//								} `json:"vssLoggingContext"`
//							} `json:"loggingContext"`
//							WatchEndpointSupportedOnesieConfig struct {
//								HTML5PlaybackOnesieConfig struct {
//									CommonConfig struct {
//										URL string `json:"url"`
//									} `json:"commonConfig"`
//								} `json:"html5PlaybackOnesieConfig"`
//							} `json:"watchEndpointSupportedOnesieConfig"`
//						} `json:"watchEndpoint"`
//					} `json:"navigationEndpoint"`
//					TrackingParams string `json:"trackingParams"`
//				} `json:"buttonRenderer"`
//			} `json:"shufflePlayButton"`
//			OnDescriptionTap struct {
//				ClickTrackingParams string `json:"clickTrackingParams"`
//				OpenPopupAction     struct {
//					Popup struct {
//						FancyDismissibleDialogRenderer struct {
//							DialogMessage struct {
//								Runs []struct {
//									Text string `json:"text"`
//								} `json:"runs"`
//							} `json:"dialogMessage"`
//							Title struct {
//								Runs []struct {
//									Text string `json:"text"`
//								} `json:"runs"`
//							} `json:"title"`
//							ConfirmLabel struct {
//								Runs []struct {
//									Text string `json:"text"`
//								} `json:"runs"`
//							} `json:"confirmLabel"`
//							TrackingParams string `json:"trackingParams"`
//						} `json:"fancyDismissibleDialogRenderer"`
//					} `json:"popup"`
//					PopupType string `json:"popupType"`
//				} `json:"openPopupAction"`
//			} `json:"onDescriptionTap"`
//			CinematicContainer struct {
//				CinematicContainerRenderer struct {
//					BackgroundImageConfig struct {
//						Thumbnail struct {
//							Thumbnails []struct {
//								URL    string `json:"url"`
//								Width  int    `json:"width"`
//								Height int    `json:"height"`
//							} `json:"thumbnails"`
//						} `json:"thumbnail"`
//					} `json:"backgroundImageConfig"`
//					GradientColorConfig []struct {
//						LightThemeColor int64 `json:"lightThemeColor"`
//						DarkThemeColor  int64 `json:"darkThemeColor"`
//						StartLocation   int   `json:"startLocation"`
//					} `json:"gradientColorConfig"`
//					Config struct {
//						LightThemeBackgroundColor int64 `json:"lightThemeBackgroundColor"`
//						DarkThemeBackgroundColor  int64 `json:"darkThemeBackgroundColor"`
//						ColorSourceSizeMultiplier int   `json:"colorSourceSizeMultiplier"`
//						ApplyClientImageBlur      bool  `json:"applyClientImageBlur"`
//					} `json:"config"`
//				} `json:"cinematicContainerRenderer"`
//			} `json:"cinematicContainer"`
//			Byline []struct {
//				PlaylistBylineRenderer struct {
//					Text struct {
//						Runs []struct {
//							Text string `json:"text"`
//						} `json:"runs"`
//					} `json:"text"`
//				} `json:"playlistBylineRenderer"`
//			} `json:"byline"`
//			DescriptionTapText struct {
//				Runs []struct {
//					Text string `json:"text"`
//				} `json:"runs"`
//			} `json:"descriptionTapText"`
//		} `json:"playlistHeaderRenderer"`
//	} `json:"header"`
//	Alerts []struct {
//		AlertWithButtonRenderer struct {
//			Type string `json:"type"`
//			Text struct {
//				SimpleText string `json:"simpleText"`
//			} `json:"text"`
//			DismissButton struct {
//				ButtonRenderer struct {
//					Style      string `json:"style"`
//					Size       string `json:"size"`
//					IsDisabled bool   `json:"isDisabled"`
//					Icon       struct {
//						IconType string `json:"iconType"`
//					} `json:"icon"`
//					TrackingParams    string `json:"trackingParams"`
//					AccessibilityData struct {
//						AccessibilityData struct {
//							Label string `json:"label"`
//						} `json:"accessibilityData"`
//					} `json:"accessibilityData"`
//				} `json:"buttonRenderer"`
//			} `json:"dismissButton"`
//		} `json:"alertWithButtonRenderer"`
//	} `json:"alerts"`
//	Metadata struct {
//		PlaylistMetadataRenderer struct {
//			Title                  string `json:"title"`
//			Description            string `json:"description"`
//			AndroidAppindexingLink string `json:"androidAppindexingLink"`
//			IosAppindexingLink     string `json:"iosAppindexingLink"`
//		} `json:"playlistMetadataRenderer"`
//	} `json:"metadata"`
//	TrackingParams string `json:"trackingParams"`
//	Topbar         struct {
//		DesktopTopbarRenderer struct {
//			Logo struct {
//				TopbarLogoRenderer struct {
//					IconImage struct {
//						IconType string `json:"iconType"`
//					} `json:"iconImage"`
//					TooltipText struct {
//						Runs []struct {
//							Text string `json:"text"`
//						} `json:"runs"`
//					} `json:"tooltipText"`
//					Endpoint struct {
//						ClickTrackingParams string `json:"clickTrackingParams"`
//						CommandMetadata     struct {
//							WebCommandMetadata struct {
//								URL         string `json:"url"`
//								WebPageType string `json:"webPageType"`
//								RootVe      int    `json:"rootVe"`
//								APIURL      string `json:"apiUrl"`
//							} `json:"webCommandMetadata"`
//						} `json:"commandMetadata"`
//						BrowseEndpoint struct {
//							BrowseID string `json:"browseId"`
//						} `json:"browseEndpoint"`
//					} `json:"endpoint"`
//					TrackingParams    string `json:"trackingParams"`
//					OverrideEntityKey string `json:"overrideEntityKey"`
//				} `json:"topbarLogoRenderer"`
//			} `json:"logo"`
//			Searchbox struct {
//				FusionSearchboxRenderer struct {
//					Icon struct {
//						IconType string `json:"iconType"`
//					} `json:"icon"`
//					PlaceholderText struct {
//						Runs []struct {
//							Text string `json:"text"`
//						} `json:"runs"`
//					} `json:"placeholderText"`
//					Config struct {
//						WebSearchboxConfig struct {
//							RequestLanguage     string `json:"requestLanguage"`
//							RequestDomain       string `json:"requestDomain"`
//							HasOnscreenKeyboard bool   `json:"hasOnscreenKeyboard"`
//							FocusSearchbox      bool   `json:"focusSearchbox"`
//						} `json:"webSearchboxConfig"`
//					} `json:"config"`
//					TrackingParams string `json:"trackingParams"`
//					SearchEndpoint struct {
//						ClickTrackingParams string `json:"clickTrackingParams"`
//						CommandMetadata     struct {
//							WebCommandMetadata struct {
//								URL         string `json:"url"`
//								WebPageType string `json:"webPageType"`
//								RootVe      int    `json:"rootVe"`
//							} `json:"webCommandMetadata"`
//						} `json:"commandMetadata"`
//						SearchEndpoint struct {
//							Query string `json:"query"`
//						} `json:"searchEndpoint"`
//					} `json:"searchEndpoint"`
//					ClearButton struct {
//						ButtonRenderer struct {
//							Style      string `json:"style"`
//							Size       string `json:"size"`
//							IsDisabled bool   `json:"isDisabled"`
//							Icon       struct {
//								IconType string `json:"iconType"`
//							} `json:"icon"`
//							TrackingParams    string `json:"trackingParams"`
//							AccessibilityData struct {
//								AccessibilityData struct {
//									Label string `json:"label"`
//								} `json:"accessibilityData"`
//							} `json:"accessibilityData"`
//						} `json:"buttonRenderer"`
//					} `json:"clearButton"`
//				} `json:"fusionSearchboxRenderer"`
//			} `json:"searchbox"`
//			TrackingParams string `json:"trackingParams"`
//			CountryCode    string `json:"countryCode"`
//			TopbarButtons  []struct {
//				TopbarMenuButtonRenderer struct {
//					Icon struct {
//						IconType string `json:"iconType"`
//					} `json:"icon"`
//					MenuRequest struct {
//						ClickTrackingParams string `json:"clickTrackingParams"`
//						CommandMetadata     struct {
//							WebCommandMetadata struct {
//								SendPost bool   `json:"sendPost"`
//								APIURL   string `json:"apiUrl"`
//							} `json:"webCommandMetadata"`
//						} `json:"commandMetadata"`
//						SignalServiceEndpoint struct {
//							Signal  string `json:"signal"`
//							Actions []struct {
//								ClickTrackingParams string `json:"clickTrackingParams"`
//								OpenPopupAction     struct {
//									Popup struct {
//										MultiPageMenuRenderer struct {
//											TrackingParams     string `json:"trackingParams"`
//											Style              string `json:"style"`
//											ShowLoadingSpinner bool   `json:"showLoadingSpinner"`
//										} `json:"multiPageMenuRenderer"`
//									} `json:"popup"`
//									PopupType string `json:"popupType"`
//									BeReused  bool   `json:"beReused"`
//								} `json:"openPopupAction"`
//							} `json:"actions"`
//						} `json:"signalServiceEndpoint"`
//					} `json:"menuRequest"`
//					TrackingParams string `json:"trackingParams"`
//					Accessibility  struct {
//						AccessibilityData struct {
//							Label string `json:"label"`
//						} `json:"accessibilityData"`
//					} `json:"accessibility"`
//					Tooltip string `json:"tooltip"`
//					Style   string `json:"style"`
//				} `json:"topbarMenuButtonRenderer,omitempty"`
//				ButtonRenderer struct {
//					Style string `json:"style"`
//					Size  string `json:"size"`
//					Text  struct {
//						Runs []struct {
//							Text string `json:"text"`
//						} `json:"runs"`
//					} `json:"text"`
//					Icon struct {
//						IconType string `json:"iconType"`
//					} `json:"icon"`
//					NavigationEndpoint struct {
//						ClickTrackingParams string `json:"clickTrackingParams"`
//						CommandMetadata     struct {
//							WebCommandMetadata struct {
//								URL         string `json:"url"`
//								WebPageType string `json:"webPageType"`
//								RootVe      int    `json:"rootVe"`
//							} `json:"webCommandMetadata"`
//						} `json:"commandMetadata"`
//						SignInEndpoint struct {
//							IdamTag string `json:"idamTag"`
//						} `json:"signInEndpoint"`
//					} `json:"navigationEndpoint"`
//					TrackingParams string `json:"trackingParams"`
//					TargetID       string `json:"targetId"`
//				} `json:"buttonRenderer,omitempty"`
//			} `json:"topbarButtons"`
//			HotkeyDialog struct {
//				HotkeyDialogRenderer struct {
//					Title struct {
//						Runs []struct {
//							Text string `json:"text"`
//						} `json:"runs"`
//					} `json:"title"`
//					Sections []struct {
//						HotkeyDialogSectionRenderer struct {
//							Title struct {
//								Runs []struct {
//									Text string `json:"text"`
//								} `json:"runs"`
//							} `json:"title"`
//							Options []struct {
//								HotkeyDialogSectionOptionRenderer struct {
//									Label struct {
//										Runs []struct {
//											Text string `json:"text"`
//										} `json:"runs"`
//									} `json:"label"`
//									Hotkey string `json:"hotkey"`
//								} `json:"hotkeyDialogSectionOptionRenderer,omitempty"`
//								HotkeyDialogSectionOptionRenderer0 struct {
//									Label struct {
//										Runs []struct {
//											Text string `json:"text"`
//										} `json:"runs"`
//									} `json:"label"`
//									Hotkey                   string `json:"hotkey"`
//									HotkeyAccessibilityLabel struct {
//										AccessibilityData struct {
//											Label string `json:"label"`
//										} `json:"accessibilityData"`
//									} `json:"hotkeyAccessibilityLabel"`
//								} `json:"hotkeyDialogSectionOptionRenderer,omitempty"`
//								HotkeyDialogSectionOptionRenderer1 struct {
//									Label struct {
//										Runs []struct {
//											Text string `json:"text"`
//										} `json:"runs"`
//									} `json:"label"`
//									Hotkey                   string `json:"hotkey"`
//									HotkeyAccessibilityLabel struct {
//										AccessibilityData struct {
//											Label string `json:"label"`
//										} `json:"accessibilityData"`
//									} `json:"hotkeyAccessibilityLabel"`
//								} `json:"hotkeyDialogSectionOptionRenderer,omitempty"`
//								HotkeyDialogSectionOptionRenderer2 struct {
//									Label struct {
//										Runs []struct {
//											Text string `json:"text"`
//										} `json:"runs"`
//									} `json:"label"`
//									Hotkey                   string `json:"hotkey"`
//									HotkeyAccessibilityLabel struct {
//										AccessibilityData struct {
//											Label string `json:"label"`
//										} `json:"accessibilityData"`
//									} `json:"hotkeyAccessibilityLabel"`
//								} `json:"hotkeyDialogSectionOptionRenderer,omitempty"`
//								HotkeyDialogSectionOptionRenderer3 struct {
//									Label struct {
//										Runs []struct {
//											Text string `json:"text"`
//										} `json:"runs"`
//									} `json:"label"`
//									Hotkey                   string `json:"hotkey"`
//									HotkeyAccessibilityLabel struct {
//										AccessibilityData struct {
//											Label string `json:"label"`
//										} `json:"accessibilityData"`
//									} `json:"hotkeyAccessibilityLabel"`
//								} `json:"hotkeyDialogSectionOptionRenderer,omitempty"`
//							} `json:"options"`
//						} `json:"hotkeyDialogSectionRenderer"`
//					} `json:"sections"`
//					DismissButton struct {
//						ButtonRenderer struct {
//							Style      string `json:"style"`
//							Size       string `json:"size"`
//							IsDisabled bool   `json:"isDisabled"`
//							Text       struct {
//								Runs []struct {
//									Text string `json:"text"`
//								} `json:"runs"`
//							} `json:"text"`
//							TrackingParams string `json:"trackingParams"`
//						} `json:"buttonRenderer"`
//					} `json:"dismissButton"`
//					TrackingParams string `json:"trackingParams"`
//				} `json:"hotkeyDialogRenderer"`
//			} `json:"hotkeyDialog"`
//			BackButton struct {
//				ButtonRenderer struct {
//					TrackingParams string `json:"trackingParams"`
//					Command        struct {
//						ClickTrackingParams string `json:"clickTrackingParams"`
//						CommandMetadata     struct {
//							WebCommandMetadata struct {
//								SendPost bool `json:"sendPost"`
//							} `json:"webCommandMetadata"`
//						} `json:"commandMetadata"`
//						SignalServiceEndpoint struct {
//							Signal  string `json:"signal"`
//							Actions []struct {
//								ClickTrackingParams string `json:"clickTrackingParams"`
//								SignalAction        struct {
//									Signal string `json:"signal"`
//								} `json:"signalAction"`
//							} `json:"actions"`
//						} `json:"signalServiceEndpoint"`
//					} `json:"command"`
//				} `json:"buttonRenderer"`
//			} `json:"backButton"`
//			ForwardButton struct {
//				ButtonRenderer struct {
//					TrackingParams string `json:"trackingParams"`
//					Command        struct {
//						ClickTrackingParams string `json:"clickTrackingParams"`
//						CommandMetadata     struct {
//							WebCommandMetadata struct {
//								SendPost bool `json:"sendPost"`
//							} `json:"webCommandMetadata"`
//						} `json:"commandMetadata"`
//						SignalServiceEndpoint struct {
//							Signal  string `json:"signal"`
//							Actions []struct {
//								ClickTrackingParams string `json:"clickTrackingParams"`
//								SignalAction        struct {
//									Signal string `json:"signal"`
//								} `json:"signalAction"`
//							} `json:"actions"`
//						} `json:"signalServiceEndpoint"`
//					} `json:"command"`
//				} `json:"buttonRenderer"`
//			} `json:"forwardButton"`
//			A11YSkipNavigationButton struct {
//				ButtonRenderer struct {
//					Style      string `json:"style"`
//					Size       string `json:"size"`
//					IsDisabled bool   `json:"isDisabled"`
//					Text       struct {
//						Runs []struct {
//							Text string `json:"text"`
//						} `json:"runs"`
//					} `json:"text"`
//					TrackingParams string `json:"trackingParams"`
//					Command        struct {
//						ClickTrackingParams string `json:"clickTrackingParams"`
//						CommandMetadata     struct {
//							WebCommandMetadata struct {
//								SendPost bool `json:"sendPost"`
//							} `json:"webCommandMetadata"`
//						} `json:"commandMetadata"`
//						SignalServiceEndpoint struct {
//							Signal  string `json:"signal"`
//							Actions []struct {
//								ClickTrackingParams string `json:"clickTrackingParams"`
//								SignalAction        struct {
//									Signal string `json:"signal"`
//								} `json:"signalAction"`
//							} `json:"actions"`
//						} `json:"signalServiceEndpoint"`
//					} `json:"command"`
//				} `json:"buttonRenderer"`
//			} `json:"a11ySkipNavigationButton"`
//		} `json:"desktopTopbarRenderer"`
//	} `json:"topbar"`
//	Microformat struct {
//		MicroformatDataRenderer struct {
//			URLCanonical string `json:"urlCanonical"`
//			Title        string `json:"title"`
//			Description  string `json:"description"`
//			Thumbnail    struct {
//				Thumbnails []struct {
//					URL    string `json:"url"`
//					Width  int    `json:"width"`
//					Height int    `json:"height"`
//				} `json:"thumbnails"`
//			} `json:"thumbnail"`
//			SiteName           string `json:"siteName"`
//			AppName            string `json:"appName"`
//			AndroidPackage     string `json:"androidPackage"`
//			IosAppStoreID      string `json:"iosAppStoreId"`
//			IosAppArguments    string `json:"iosAppArguments"`
//			OgType             string `json:"ogType"`
//			URLApplinksWeb     string `json:"urlApplinksWeb"`
//			URLApplinksIos     string `json:"urlApplinksIos"`
//			URLApplinksAndroid string `json:"urlApplinksAndroid"`
//			URLTwitterIos      string `json:"urlTwitterIos"`
//			URLTwitterAndroid  string `json:"urlTwitterAndroid"`
//			TwitterCardType    string `json:"twitterCardType"`
//			TwitterSiteHandle  string `json:"twitterSiteHandle"`
//			SchemaDotOrgType   string `json:"schemaDotOrgType"`
//			Noindex            bool   `json:"noindex"`
//			Unlisted           bool   `json:"unlisted"`
//			LinkAlternates     []struct {
//				HrefURL string `json:"hrefUrl"`
//			} `json:"linkAlternates"`
//		} `json:"microformatDataRenderer"`
//	} `json:"microformat"`
//	Sidebar struct {
//		PlaylistSidebarRenderer struct {
//			Items []struct {
//				PlaylistSidebarPrimaryInfoRenderer struct {
//					ThumbnailRenderer struct {
//						PlaylistVideoThumbnailRenderer struct {
//							Thumbnail struct {
//								Thumbnails []struct {
//									URL    string `json:"url"`
//									Width  int    `json:"width"`
//									Height int    `json:"height"`
//								} `json:"thumbnails"`
//							} `json:"thumbnail"`
//							TrackingParams string `json:"trackingParams"`
//						} `json:"playlistVideoThumbnailRenderer"`
//					} `json:"thumbnailRenderer"`
//					Title struct {
//						Runs []struct {
//							Text               string `json:"text"`
//							NavigationEndpoint struct {
//								ClickTrackingParams string `json:"clickTrackingParams"`
//								CommandMetadata     struct {
//									WebCommandMetadata struct {
//										URL         string `json:"url"`
//										WebPageType string `json:"webPageType"`
//										RootVe      int    `json:"rootVe"`
//									} `json:"webCommandMetadata"`
//								} `json:"commandMetadata"`
//								WatchEndpoint struct {
//									VideoID        string `json:"videoId"`
//									PlaylistID     string `json:"playlistId"`
//									PlayerParams   string `json:"playerParams"`
//									LoggingContext struct {
//										VssLoggingContext struct {
//											SerializedContextData string `json:"serializedContextData"`
//										} `json:"vssLoggingContext"`
//									} `json:"loggingContext"`
//									WatchEndpointSupportedOnesieConfig struct {
//										HTML5PlaybackOnesieConfig struct {
//											CommonConfig struct {
//												URL string `json:"url"`
//											} `json:"commonConfig"`
//										} `json:"html5PlaybackOnesieConfig"`
//									} `json:"watchEndpointSupportedOnesieConfig"`
//								} `json:"watchEndpoint"`
//							} `json:"navigationEndpoint"`
//						} `json:"runs"`
//					} `json:"title"`
//					Stats []struct {
//						Runs []struct {
//							Text string `json:"text"`
//						} `json:"runs,omitempty"`
//						SimpleText string `json:"simpleText,omitempty"`
//					} `json:"stats"`
//					Menu struct {
//						MenuRenderer struct {
//							Items []struct {
//								MenuNavigationItemRenderer struct {
//									Text struct {
//										SimpleText string `json:"simpleText"`
//									} `json:"text"`
//									Icon struct {
//										IconType string `json:"iconType"`
//									} `json:"icon"`
//									NavigationEndpoint struct {
//										ClickTrackingParams string `json:"clickTrackingParams"`
//										CommandMetadata     struct {
//											WebCommandMetadata struct {
//												URL         string `json:"url"`
//												WebPageType string `json:"webPageType"`
//												RootVe      int    `json:"rootVe"`
//												APIURL      string `json:"apiUrl"`
//											} `json:"webCommandMetadata"`
//										} `json:"commandMetadata"`
//										BrowseEndpoint struct {
//											BrowseID       string `json:"browseId"`
//											Params         string `json:"params"`
//											Nofollow       bool   `json:"nofollow"`
//											NavigationType string `json:"navigationType"`
//										} `json:"browseEndpoint"`
//									} `json:"navigationEndpoint"`
//									TrackingParams string `json:"trackingParams"`
//								} `json:"menuNavigationItemRenderer"`
//							} `json:"items"`
//							TrackingParams  string `json:"trackingParams"`
//							TopLevelButtons []struct {
//								ToggleButtonRenderer struct {
//									Style struct {
//										StyleType string `json:"styleType"`
//									} `json:"style"`
//									Size struct {
//										SizeType string `json:"sizeType"`
//									} `json:"size"`
//									IsToggled   bool `json:"isToggled"`
//									IsDisabled  bool `json:"isDisabled"`
//									DefaultIcon struct {
//										IconType string `json:"iconType"`
//									} `json:"defaultIcon"`
//									ToggledIcon struct {
//										IconType string `json:"iconType"`
//									} `json:"toggledIcon"`
//									TrackingParams            string `json:"trackingParams"`
//									DefaultTooltip            string `json:"defaultTooltip"`
//									ToggledTooltip            string `json:"toggledTooltip"`
//									DefaultNavigationEndpoint struct {
//										ClickTrackingParams string `json:"clickTrackingParams"`
//										CommandMetadata     struct {
//											WebCommandMetadata struct {
//												IgnoreNavigation bool `json:"ignoreNavigation"`
//											} `json:"webCommandMetadata"`
//										} `json:"commandMetadata"`
//										ModalEndpoint struct {
//											Modal struct {
//												ModalWithTitleAndButtonRenderer struct {
//													Title struct {
//														SimpleText string `json:"simpleText"`
//													} `json:"title"`
//													Content struct {
//														SimpleText string `json:"simpleText"`
//													} `json:"content"`
//													Button struct {
//														ButtonRenderer struct {
//															Style      string `json:"style"`
//															Size       string `json:"size"`
//															IsDisabled bool   `json:"isDisabled"`
//															Text       struct {
//																SimpleText string `json:"simpleText"`
//															} `json:"text"`
//															NavigationEndpoint struct {
//																ClickTrackingParams string `json:"clickTrackingParams"`
//																CommandMetadata     struct {
//																	WebCommandMetadata struct {
//																		URL         string `json:"url"`
//																		WebPageType string `json:"webPageType"`
//																		RootVe      int    `json:"rootVe"`
//																	} `json:"webCommandMetadata"`
//																} `json:"commandMetadata"`
//																SignInEndpoint struct {
//																	NextEndpoint struct {
//																		ClickTrackingParams string `json:"clickTrackingParams"`
//																		CommandMetadata     struct {
//																			WebCommandMetadata struct {
//																				URL         string `json:"url"`
//																				WebPageType string `json:"webPageType"`
//																				RootVe      int    `json:"rootVe"`
//																				APIURL      string `json:"apiUrl"`
//																			} `json:"webCommandMetadata"`
//																		} `json:"commandMetadata"`
//																		BrowseEndpoint struct {
//																			BrowseID string `json:"browseId"`
//																		} `json:"browseEndpoint"`
//																	} `json:"nextEndpoint"`
//																	IdamTag string `json:"idamTag"`
//																} `json:"signInEndpoint"`
//															} `json:"navigationEndpoint"`
//															TrackingParams string `json:"trackingParams"`
//														} `json:"buttonRenderer"`
//													} `json:"button"`
//												} `json:"modalWithTitleAndButtonRenderer"`
//											} `json:"modal"`
//										} `json:"modalEndpoint"`
//									} `json:"defaultNavigationEndpoint"`
//									AccessibilityData struct {
//										AccessibilityData struct {
//											Label string `json:"label"`
//										} `json:"accessibilityData"`
//									} `json:"accessibilityData"`
//									ToggledAccessibilityData struct {
//										AccessibilityData struct {
//											Label string `json:"label"`
//										} `json:"accessibilityData"`
//									} `json:"toggledAccessibilityData"`
//								} `json:"toggleButtonRenderer,omitempty"`
//								ButtonRenderer struct {
//									Style      string `json:"style"`
//									Size       string `json:"size"`
//									IsDisabled bool   `json:"isDisabled"`
//									Icon       struct {
//										IconType string `json:"iconType"`
//									} `json:"icon"`
//									NavigationEndpoint struct {
//										ClickTrackingParams string `json:"clickTrackingParams"`
//										CommandMetadata     struct {
//											WebCommandMetadata struct {
//												URL         string `json:"url"`
//												WebPageType string `json:"webPageType"`
//												RootVe      int    `json:"rootVe"`
//											} `json:"webCommandMetadata"`
//										} `json:"commandMetadata"`
//										WatchEndpoint struct {
//											VideoID        string `json:"videoId"`
//											PlaylistID     string `json:"playlistId"`
//											Params         string `json:"params"`
//											PlayerParams   string `json:"playerParams"`
//											LoggingContext struct {
//												VssLoggingContext struct {
//													SerializedContextData string `json:"serializedContextData"`
//												} `json:"vssLoggingContext"`
//											} `json:"loggingContext"`
//											WatchEndpointSupportedOnesieConfig struct {
//												HTML5PlaybackOnesieConfig struct {
//													CommonConfig struct {
//														URL string `json:"url"`
//													} `json:"commonConfig"`
//												} `json:"html5PlaybackOnesieConfig"`
//											} `json:"watchEndpointSupportedOnesieConfig"`
//										} `json:"watchEndpoint"`
//									} `json:"navigationEndpoint"`
//									Accessibility struct {
//										Label string `json:"label"`
//									} `json:"accessibility"`
//									Tooltip        string `json:"tooltip"`
//									TrackingParams string `json:"trackingParams"`
//								} `json:"buttonRenderer,omitempty"`
//								ButtonRenderer0 struct {
//									Style           string `json:"style"`
//									Size            string `json:"size"`
//									IsDisabled      bool   `json:"isDisabled"`
//									ServiceEndpoint struct {
//										ClickTrackingParams string `json:"clickTrackingParams"`
//										CommandMetadata     struct {
//											WebCommandMetadata struct {
//												SendPost bool   `json:"sendPost"`
//												APIURL   string `json:"apiUrl"`
//											} `json:"webCommandMetadata"`
//										} `json:"commandMetadata"`
//										ShareEntityServiceEndpoint struct {
//											SerializedShareEntity string `json:"serializedShareEntity"`
//											Commands              []struct {
//												ClickTrackingParams string `json:"clickTrackingParams"`
//												OpenPopupAction     struct {
//													Popup struct {
//														UnifiedSharePanelRenderer struct {
//															TrackingParams     string `json:"trackingParams"`
//															ShowLoadingSpinner bool   `json:"showLoadingSpinner"`
//														} `json:"unifiedSharePanelRenderer"`
//													} `json:"popup"`
//													PopupType string `json:"popupType"`
//													BeReused  bool   `json:"beReused"`
//												} `json:"openPopupAction"`
//											} `json:"commands"`
//										} `json:"shareEntityServiceEndpoint"`
//									} `json:"serviceEndpoint"`
//									Icon struct {
//										IconType string `json:"iconType"`
//									} `json:"icon"`
//									Accessibility struct {
//										Label string `json:"label"`
//									} `json:"accessibility"`
//									Tooltip        string `json:"tooltip"`
//									TrackingParams string `json:"trackingParams"`
//								} `json:"buttonRenderer,omitempty"`
//							} `json:"topLevelButtons"`
//							Accessibility struct {
//								AccessibilityData struct {
//									Label string `json:"label"`
//								} `json:"accessibilityData"`
//							} `json:"accessibility"`
//							TargetID string `json:"targetId"`
//						} `json:"menuRenderer"`
//					} `json:"menu"`
//					ThumbnailOverlays []struct {
//						ThumbnailOverlaySidePanelRenderer struct {
//							Text struct {
//								SimpleText string `json:"simpleText"`
//							} `json:"text"`
//							Icon struct {
//								IconType string `json:"iconType"`
//							} `json:"icon"`
//						} `json:"thumbnailOverlaySidePanelRenderer"`
//					} `json:"thumbnailOverlays"`
//					NavigationEndpoint struct {
//						ClickTrackingParams string `json:"clickTrackingParams"`
//						CommandMetadata     struct {
//							WebCommandMetadata struct {
//								URL         string `json:"url"`
//								WebPageType string `json:"webPageType"`
//								RootVe      int    `json:"rootVe"`
//							} `json:"webCommandMetadata"`
//						} `json:"commandMetadata"`
//						WatchEndpoint struct {
//							VideoID        string `json:"videoId"`
//							PlaylistID     string `json:"playlistId"`
//							PlayerParams   string `json:"playerParams"`
//							LoggingContext struct {
//								VssLoggingContext struct {
//									SerializedContextData string `json:"serializedContextData"`
//								} `json:"vssLoggingContext"`
//							} `json:"loggingContext"`
//							WatchEndpointSupportedOnesieConfig struct {
//								HTML5PlaybackOnesieConfig struct {
//									CommonConfig struct {
//										URL string `json:"url"`
//									} `json:"commonConfig"`
//								} `json:"html5PlaybackOnesieConfig"`
//							} `json:"watchEndpointSupportedOnesieConfig"`
//						} `json:"watchEndpoint"`
//					} `json:"navigationEndpoint"`
//					Description struct {
//						SimpleText string `json:"simpleText"`
//					} `json:"description"`
//					ShowMoreText struct {
//						Runs []struct {
//							Text string `json:"text"`
//						} `json:"runs"`
//					} `json:"showMoreText"`
//				} `json:"playlistSidebarPrimaryInfoRenderer,omitempty"`
//				PlaylistSidebarSecondaryInfoRenderer struct {
//					VideoOwner struct {
//						VideoOwnerRenderer struct {
//							Thumbnail struct {
//								Thumbnails []struct {
//									URL    string `json:"url"`
//									Width  int    `json:"width"`
//									Height int    `json:"height"`
//								} `json:"thumbnails"`
//							} `json:"thumbnail"`
//							Title struct {
//								Runs []struct {
//									Text               string `json:"text"`
//									NavigationEndpoint struct {
//										ClickTrackingParams string `json:"clickTrackingParams"`
//										CommandMetadata     struct {
//											WebCommandMetadata struct {
//												URL         string `json:"url"`
//												WebPageType string `json:"webPageType"`
//												RootVe      int    `json:"rootVe"`
//												APIURL      string `json:"apiUrl"`
//											} `json:"webCommandMetadata"`
//										} `json:"commandMetadata"`
//										BrowseEndpoint struct {
//											BrowseID         string `json:"browseId"`
//											CanonicalBaseURL string `json:"canonicalBaseUrl"`
//										} `json:"browseEndpoint"`
//									} `json:"navigationEndpoint"`
//								} `json:"runs"`
//							} `json:"title"`
//							NavigationEndpoint struct {
//								ClickTrackingParams string `json:"clickTrackingParams"`
//								CommandMetadata     struct {
//									WebCommandMetadata struct {
//										URL         string `json:"url"`
//										WebPageType string `json:"webPageType"`
//										RootVe      int    `json:"rootVe"`
//										APIURL      string `json:"apiUrl"`
//									} `json:"webCommandMetadata"`
//								} `json:"commandMetadata"`
//								BrowseEndpoint struct {
//									BrowseID         string `json:"browseId"`
//									CanonicalBaseURL string `json:"canonicalBaseUrl"`
//								} `json:"browseEndpoint"`
//							} `json:"navigationEndpoint"`
//							TrackingParams string `json:"trackingParams"`
//						} `json:"videoOwnerRenderer"`
//					} `json:"videoOwner"`
//					Button struct {
//						ButtonRenderer struct {
//							Style      string `json:"style"`
//							Size       string `json:"size"`
//							IsDisabled bool   `json:"isDisabled"`
//							Text       struct {
//								Runs []struct {
//									Text string `json:"text"`
//								} `json:"runs"`
//							} `json:"text"`
//							NavigationEndpoint struct {
//								ClickTrackingParams string `json:"clickTrackingParams"`
//								CommandMetadata     struct {
//									WebCommandMetadata struct {
//										IgnoreNavigation bool `json:"ignoreNavigation"`
//									} `json:"webCommandMetadata"`
//								} `json:"commandMetadata"`
//								ModalEndpoint struct {
//									Modal struct {
//										ModalWithTitleAndButtonRenderer struct {
//											Title struct {
//												SimpleText string `json:"simpleText"`
//											} `json:"title"`
//											Content struct {
//												SimpleText string `json:"simpleText"`
//											} `json:"content"`
//											Button struct {
//												ButtonRenderer struct {
//													Style      string `json:"style"`
//													Size       string `json:"size"`
//													IsDisabled bool   `json:"isDisabled"`
//													Text       struct {
//														SimpleText string `json:"simpleText"`
//													} `json:"text"`
//													NavigationEndpoint struct {
//														ClickTrackingParams string `json:"clickTrackingParams"`
//														CommandMetadata     struct {
//															WebCommandMetadata struct {
//																URL         string `json:"url"`
//																WebPageType string `json:"webPageType"`
//																RootVe      int    `json:"rootVe"`
//															} `json:"webCommandMetadata"`
//														} `json:"commandMetadata"`
//														SignInEndpoint struct {
//															NextEndpoint struct {
//																ClickTrackingParams string `json:"clickTrackingParams"`
//																CommandMetadata     struct {
//																	WebCommandMetadata struct {
//																		URL         string `json:"url"`
//																		WebPageType string `json:"webPageType"`
//																		RootVe      int    `json:"rootVe"`
//																		APIURL      string `json:"apiUrl"`
//																	} `json:"webCommandMetadata"`
//																} `json:"commandMetadata"`
//																BrowseEndpoint struct {
//																	BrowseID string `json:"browseId"`
//																} `json:"browseEndpoint"`
//															} `json:"nextEndpoint"`
//															ContinueAction string `json:"continueAction"`
//															IdamTag        string `json:"idamTag"`
//														} `json:"signInEndpoint"`
//													} `json:"navigationEndpoint"`
//													TrackingParams string `json:"trackingParams"`
//												} `json:"buttonRenderer"`
//											} `json:"button"`
//										} `json:"modalWithTitleAndButtonRenderer"`
//									} `json:"modal"`
//								} `json:"modalEndpoint"`
//							} `json:"navigationEndpoint"`
//							TrackingParams string `json:"trackingParams"`
//						} `json:"buttonRenderer"`
//					} `json:"button"`
//				} `json:"playlistSidebarSecondaryInfoRenderer,omitempty"`
//			} `json:"items"`
//			TrackingParams string `json:"trackingParams"`
//		} `json:"playlistSidebarRenderer"`
//	} `json:"sidebar"`
//}

//type PlayListStruct2 struct {
//	ResponseContext struct {
//		ServiceTrackingParams []struct {
//			Service string `json:"service"`
//			Params  []struct {
//				Key   string `json:"key"`
//				Value string `json:"value"`
//			} `json:"params"`
//		} `json:"serviceTrackingParams"`
//		MaxAgeSeconds             int `json:"maxAgeSeconds"`
//		MainAppWebResponseContext struct {
//			LoggedOut     bool   `json:"loggedOut"`
//			TrackingParam string `json:"trackingParam"`
//		} `json:"mainAppWebResponseContext"`
//		WebResponseContextExtensionData struct {
//			HasDecorated bool `json:"hasDecorated"`
//		} `json:"webResponseContextExtensionData"`
//	} `json:"responseContext"`
//	PlayabilityStatus struct {
//		Status          string `json:"status"`
//		PlayableInEmbed bool   `json:"playableInEmbed"`
//		Miniplayer      struct {
//			MiniplayerRenderer struct {
//				PlaybackMode string `json:"playbackMode"`
//			} `json:"miniplayerRenderer"`
//		} `json:"miniplayer"`
//		ContextParams string `json:"contextParams"`
//	} `json:"playabilityStatus"`
//	StreamingData struct {
//		ExpiresInSeconds string `json:"expiresInSeconds"`
//		Formats          []struct {
//			Itag             int    `json:"itag"`
//			URL              string `json:"url"`
//			MimeType         string `json:"mimeType"`
//			Bitrate          int    `json:"bitrate"`
//			Width            int    `json:"width"`
//			Height           int    `json:"height"`
//			LastModified     string `json:"lastModified"`
//			Quality          string `json:"quality"`
//			Fps              int    `json:"fps"`
//			QualityLabel     string `json:"qualityLabel"`
//			ProjectionType   string `json:"projectionType"`
//			AudioQuality     string `json:"audioQuality"`
//			ApproxDurationMs string `json:"approxDurationMs"`
//			AudioSampleRate  string `json:"audioSampleRate"`
//			AudioChannels    int    `json:"audioChannels"`
//		} `json:"formats"`
//		AdaptiveFormats []struct {
//			Itag      int    `json:"itag"`
//			URL       string `json:"url"`
//			MimeType  string `json:"mimeType"`
//			Bitrate   int    `json:"bitrate"`
//			Width     int    `json:"width,omitempty"`
//			Height    int    `json:"height,omitempty"`
//			InitRange struct {
//				Start string `json:"start"`
//				End   string `json:"end"`
//			} `json:"initRange"`
//			IndexRange struct {
//				Start string `json:"start"`
//				End   string `json:"end"`
//			} `json:"indexRange"`
//			LastModified     string `json:"lastModified"`
//			ContentLength    string `json:"contentLength"`
//			Quality          string `json:"quality"`
//			Fps              int    `json:"fps,omitempty"`
//			QualityLabel     string `json:"qualityLabel,omitempty"`
//			ProjectionType   string `json:"projectionType"`
//			AverageBitrate   int    `json:"averageBitrate"`
//			ApproxDurationMs string `json:"approxDurationMs"`
//			ColorInfo        struct {
//				Primaries               string `json:"primaries"`
//				TransferCharacteristics string `json:"transferCharacteristics"`
//				MatrixCoefficients      string `json:"matrixCoefficients"`
//			} `json:"colorInfo,omitempty"`
//			HighReplication bool    `json:"highReplication,omitempty"`
//			AudioQuality    string  `json:"audioQuality,omitempty"`
//			AudioSampleRate string  `json:"audioSampleRate,omitempty"`
//			AudioChannels   int     `json:"audioChannels,omitempty"`
//			LoudnessDb      float64 `json:"loudnessDb,omitempty"`
//		} `json:"adaptiveFormats"`
//	} `json:"streamingData"`
//	PlaybackTracking struct {
//		VideostatsPlaybackURL struct {
//			BaseURL string `json:"baseUrl"`
//		} `json:"videostatsPlaybackUrl"`
//		VideostatsDelayplayURL struct {
//			BaseURL string `json:"baseUrl"`
//		} `json:"videostatsDelayplayUrl"`
//		VideostatsWatchtimeURL struct {
//			BaseURL string `json:"baseUrl"`
//		} `json:"videostatsWatchtimeUrl"`
//		PtrackingURL struct {
//			BaseURL string `json:"baseUrl"`
//		} `json:"ptrackingUrl"`
//		QoeURL struct {
//			BaseURL string `json:"baseUrl"`
//		} `json:"qoeUrl"`
//		AtrURL struct {
//			BaseURL                 string `json:"baseUrl"`
//			ElapsedMediaTimeSeconds int    `json:"elapsedMediaTimeSeconds"`
//		} `json:"atrUrl"`
//		VideostatsScheduledFlushWalltimeSeconds []int `json:"videostatsScheduledFlushWalltimeSeconds"`
//		VideostatsDefaultFlushIntervalSeconds   int   `json:"videostatsDefaultFlushIntervalSeconds"`
//	} `json:"playbackTracking"`
//	Captions struct {
//		PlayerCaptionsTracklistRenderer struct {
//			CaptionTracks []struct {
//				BaseURL string `json:"baseUrl"`
//				Name    struct {
//					SimpleText string `json:"simpleText"`
//				} `json:"name"`
//				VssID          string `json:"vssId"`
//				LanguageCode   string `json:"languageCode"`
//				Kind           string `json:"kind"`
//				IsTranslatable bool   `json:"isTranslatable"`
//				TrackName      string `json:"trackName"`
//			} `json:"captionTracks"`
//			AudioTracks []struct {
//				CaptionTrackIndices []int `json:"captionTrackIndices"`
//			} `json:"audioTracks"`
//			TranslationLanguages []struct {
//				LanguageCode string `json:"languageCode"`
//				LanguageName struct {
//					SimpleText string `json:"simpleText"`
//				} `json:"languageName"`
//			} `json:"translationLanguages"`
//			DefaultAudioTrackIndex int `json:"defaultAudioTrackIndex"`
//		} `json:"playerCaptionsTracklistRenderer"`
//	} `json:"captions"`
//	VideoDetails struct {
//		VideoID          string   `json:"videoId"`
//		Title            string   `json:"title"`
//		LengthSeconds    string   `json:"lengthSeconds"`
//		Keywords         []string `json:"keywords"`
//		ChannelID        string   `json:"channelId"`
//		IsOwnerViewing   bool     `json:"isOwnerViewing"`
//		ShortDescription string   `json:"shortDescription"`
//		IsCrawlable      bool     `json:"isCrawlable"`
//		Thumbnail        struct {
//			Thumbnails []struct {
//				URL    string `json:"url"`
//				Width  int    `json:"width"`
//				Height int    `json:"height"`
//			} `json:"thumbnails"`
//		} `json:"thumbnail"`
//		AllowRatings           bool   `json:"allowRatings"`
//		ViewCount              string `json:"viewCount"`
//		Author                 string `json:"author"`
//		IsLowLatencyLiveStream bool   `json:"isLowLatencyLiveStream"`
//		IsPrivate              bool   `json:"isPrivate"`
//		IsUnpluggedCorpus      bool   `json:"isUnpluggedCorpus"`
//		LatencyClass           string `json:"latencyClass"`
//		IsLiveContent          bool   `json:"isLiveContent"`
//	} `json:"videoDetails"`
//	Annotations []struct {
//		PlayerAnnotationsExpandedRenderer struct {
//			FeaturedChannel struct {
//				StartTimeMs string `json:"startTimeMs"`
//				EndTimeMs   string `json:"endTimeMs"`
//				Watermark   struct {
//					Thumbnails []struct {
//						URL    string `json:"url"`
//						Width  int    `json:"width"`
//						Height int    `json:"height"`
//					} `json:"thumbnails"`
//				} `json:"watermark"`
//				TrackingParams     string `json:"trackingParams"`
//				NavigationEndpoint struct {
//					ClickTrackingParams string `json:"clickTrackingParams"`
//					CommandMetadata     struct {
//						WebCommandMetadata struct {
//							URL         string `json:"url"`
//							WebPageType string `json:"webPageType"`
//							RootVe      int    `json:"rootVe"`
//							APIURL      string `json:"apiUrl"`
//						} `json:"webCommandMetadata"`
//					} `json:"commandMetadata"`
//					BrowseEndpoint struct {
//						BrowseID string `json:"browseId"`
//					} `json:"browseEndpoint"`
//				} `json:"navigationEndpoint"`
//				ChannelName     string `json:"channelName"`
//				SubscribeButton struct {
//					SubscribeButtonRenderer struct {
//						ButtonText struct {
//							Runs []struct {
//								Text string `json:"text"`
//							} `json:"runs"`
//						} `json:"buttonText"`
//						Subscribed           bool   `json:"subscribed"`
//						Enabled              bool   `json:"enabled"`
//						Type                 string `json:"type"`
//						ChannelID            string `json:"channelId"`
//						ShowPreferences      bool   `json:"showPreferences"`
//						SubscribedButtonText struct {
//							Runs []struct {
//								Text string `json:"text"`
//							} `json:"runs"`
//						} `json:"subscribedButtonText"`
//						UnsubscribedButtonText struct {
//							Runs []struct {
//								Text string `json:"text"`
//							} `json:"runs"`
//						} `json:"unsubscribedButtonText"`
//						TrackingParams        string `json:"trackingParams"`
//						UnsubscribeButtonText struct {
//							Runs []struct {
//								Text string `json:"text"`
//							} `json:"runs"`
//						} `json:"unsubscribeButtonText"`
//						ServiceEndpoints []struct {
//							ClickTrackingParams string `json:"clickTrackingParams"`
//							CommandMetadata     struct {
//								WebCommandMetadata struct {
//									SendPost bool   `json:"sendPost"`
//									APIURL   string `json:"apiUrl"`
//								} `json:"webCommandMetadata"`
//							} `json:"commandMetadata"`
//							SubscribeEndpoint struct {
//								ChannelIds []string `json:"channelIds"`
//								Params     string   `json:"params"`
//							} `json:"subscribeEndpoint,omitempty"`
//							SignalServiceEndpoint struct {
//								Signal  string `json:"signal"`
//								Actions []struct {
//									ClickTrackingParams string `json:"clickTrackingParams"`
//									OpenPopupAction     struct {
//										Popup struct {
//											ConfirmDialogRenderer struct {
//												TrackingParams string `json:"trackingParams"`
//												DialogMessages []struct {
//													Runs []struct {
//														Text string `json:"text"`
//													} `json:"runs"`
//												} `json:"dialogMessages"`
//												ConfirmButton struct {
//													ButtonRenderer struct {
//														Style      string `json:"style"`
//														Size       string `json:"size"`
//														IsDisabled bool   `json:"isDisabled"`
//														Text       struct {
//															Runs []struct {
//																Text string `json:"text"`
//															} `json:"runs"`
//														} `json:"text"`
//														ServiceEndpoint struct {
//															ClickTrackingParams string `json:"clickTrackingParams"`
//															CommandMetadata     struct {
//																WebCommandMetadata struct {
//																	SendPost bool   `json:"sendPost"`
//																	APIURL   string `json:"apiUrl"`
//																} `json:"webCommandMetadata"`
//															} `json:"commandMetadata"`
//															UnsubscribeEndpoint struct {
//																ChannelIds []string `json:"channelIds"`
//																Params     string   `json:"params"`
//															} `json:"unsubscribeEndpoint"`
//														} `json:"serviceEndpoint"`
//														Accessibility struct {
//															Label string `json:"label"`
//														} `json:"accessibility"`
//														TrackingParams string `json:"trackingParams"`
//													} `json:"buttonRenderer"`
//												} `json:"confirmButton"`
//												CancelButton struct {
//													ButtonRenderer struct {
//														Style      string `json:"style"`
//														Size       string `json:"size"`
//														IsDisabled bool   `json:"isDisabled"`
//														Text       struct {
//															Runs []struct {
//																Text string `json:"text"`
//															} `json:"runs"`
//														} `json:"text"`
//														Accessibility struct {
//															Label string `json:"label"`
//														} `json:"accessibility"`
//														TrackingParams string `json:"trackingParams"`
//													} `json:"buttonRenderer"`
//												} `json:"cancelButton"`
//												PrimaryIsCancel bool `json:"primaryIsCancel"`
//											} `json:"confirmDialogRenderer"`
//										} `json:"popup"`
//										PopupType string `json:"popupType"`
//									} `json:"openPopupAction"`
//								} `json:"actions"`
//							} `json:"signalServiceEndpoint,omitempty"`
//						} `json:"serviceEndpoints"`
//						SubscribeAccessibility struct {
//							AccessibilityData struct {
//								Label string `json:"label"`
//							} `json:"accessibilityData"`
//						} `json:"subscribeAccessibility"`
//						UnsubscribeAccessibility struct {
//							AccessibilityData struct {
//								Label string `json:"label"`
//							} `json:"accessibilityData"`
//						} `json:"unsubscribeAccessibility"`
//						SignInEndpoint struct {
//							ClickTrackingParams string `json:"clickTrackingParams"`
//							CommandMetadata     struct {
//								WebCommandMetadata struct {
//									URL string `json:"url"`
//								} `json:"webCommandMetadata"`
//							} `json:"commandMetadata"`
//						} `json:"signInEndpoint"`
//					} `json:"subscribeButtonRenderer"`
//				} `json:"subscribeButton"`
//			} `json:"featuredChannel"`
//			AllowSwipeDismiss bool   `json:"allowSwipeDismiss"`
//			AnnotationID      string `json:"annotationId"`
//		} `json:"playerAnnotationsExpandedRenderer"`
//	} `json:"annotations"`
//	PlayerConfig struct {
//		AudioConfig struct {
//			LoudnessDb              float64 `json:"loudnessDb"`
//			PerceptualLoudnessDb    float64 `json:"perceptualLoudnessDb"`
//			EnablePerFormatLoudness bool    `json:"enablePerFormatLoudness"`
//		} `json:"audioConfig"`
//		StreamSelectionConfig struct {
//			MaxBitrate string `json:"maxBitrate"`
//		} `json:"streamSelectionConfig"`
//		MediaCommonConfig struct {
//			DynamicReadaheadConfig struct {
//				MaxReadAheadMediaTimeMs int `json:"maxReadAheadMediaTimeMs"`
//				MinReadAheadMediaTimeMs int `json:"minReadAheadMediaTimeMs"`
//				ReadAheadGrowthRateMs   int `json:"readAheadGrowthRateMs"`
//			} `json:"dynamicReadaheadConfig"`
//		} `json:"mediaCommonConfig"`
//		WebPlayerConfig struct {
//			UseCobaltTvosDash       bool `json:"useCobaltTvosDash"`
//			WebPlayerActionsPorting struct {
//				GetSharePanelCommand struct {
//					ClickTrackingParams string `json:"clickTrackingParams"`
//					CommandMetadata     struct {
//						WebCommandMetadata struct {
//							SendPost bool   `json:"sendPost"`
//							APIURL   string `json:"apiUrl"`
//						} `json:"webCommandMetadata"`
//					} `json:"commandMetadata"`
//					WebPlayerShareEntityServiceEndpoint struct {
//						SerializedShareEntity string `json:"serializedShareEntity"`
//					} `json:"webPlayerShareEntityServiceEndpoint"`
//				} `json:"getSharePanelCommand"`
//				SubscribeCommand struct {
//					ClickTrackingParams string `json:"clickTrackingParams"`
//					CommandMetadata     struct {
//						WebCommandMetadata struct {
//							SendPost bool   `json:"sendPost"`
//							APIURL   string `json:"apiUrl"`
//						} `json:"webCommandMetadata"`
//					} `json:"commandMetadata"`
//					SubscribeEndpoint struct {
//						ChannelIds []string `json:"channelIds"`
//						Params     string   `json:"params"`
//					} `json:"subscribeEndpoint"`
//				} `json:"subscribeCommand"`
//				UnsubscribeCommand struct {
//					ClickTrackingParams string `json:"clickTrackingParams"`
//					CommandMetadata     struct {
//						WebCommandMetadata struct {
//							SendPost bool   `json:"sendPost"`
//							APIURL   string `json:"apiUrl"`
//						} `json:"webCommandMetadata"`
//					} `json:"commandMetadata"`
//					UnsubscribeEndpoint struct {
//						ChannelIds []string `json:"channelIds"`
//						Params     string   `json:"params"`
//					} `json:"unsubscribeEndpoint"`
//				} `json:"unsubscribeCommand"`
//				AddToWatchLaterCommand struct {
//					ClickTrackingParams string `json:"clickTrackingParams"`
//					CommandMetadata     struct {
//						WebCommandMetadata struct {
//							SendPost bool   `json:"sendPost"`
//							APIURL   string `json:"apiUrl"`
//						} `json:"webCommandMetadata"`
//					} `json:"commandMetadata"`
//					PlaylistEditEndpoint struct {
//						PlaylistID string `json:"playlistId"`
//						Actions    []struct {
//							AddedVideoID string `json:"addedVideoId"`
//							Action       string `json:"action"`
//						} `json:"actions"`
//					} `json:"playlistEditEndpoint"`
//				} `json:"addToWatchLaterCommand"`
//				RemoveFromWatchLaterCommand struct {
//					ClickTrackingParams string `json:"clickTrackingParams"`
//					CommandMetadata     struct {
//						WebCommandMetadata struct {
//							SendPost bool   `json:"sendPost"`
//							APIURL   string `json:"apiUrl"`
//						} `json:"webCommandMetadata"`
//					} `json:"commandMetadata"`
//					PlaylistEditEndpoint struct {
//						PlaylistID string `json:"playlistId"`
//						Actions    []struct {
//							Action         string `json:"action"`
//							RemovedVideoID string `json:"removedVideoId"`
//						} `json:"actions"`
//					} `json:"playlistEditEndpoint"`
//				} `json:"removeFromWatchLaterCommand"`
//			} `json:"webPlayerActionsPorting"`
//		} `json:"webPlayerConfig"`
//	} `json:"playerConfig"`
//	Storyboards struct {
//		PlayerStoryboardSpecRenderer struct {
//			Spec             string `json:"spec"`
//			RecommendedLevel int    `json:"recommendedLevel"`
//		} `json:"playerStoryboardSpecRenderer"`
//	} `json:"storyboards"`
//	Microformat struct {
//		PlayerMicroformatRenderer struct {
//			Thumbnail struct {
//				Thumbnails []struct {
//					URL    string `json:"url"`
//					Width  int    `json:"width"`
//					Height int    `json:"height"`
//				} `json:"thumbnails"`
//			} `json:"thumbnail"`
//			Embed struct {
//				IframeURL string `json:"iframeUrl"`
//				Width     int    `json:"width"`
//				Height    int    `json:"height"`
//			} `json:"embed"`
//			Title struct {
//				SimpleText string `json:"simpleText"`
//			} `json:"title"`
//			Description struct {
//				SimpleText string `json:"simpleText"`
//			} `json:"description"`
//			LengthSeconds        string   `json:"lengthSeconds"`
//			OwnerProfileURL      string   `json:"ownerProfileUrl"`
//			ExternalChannelID    string   `json:"externalChannelId"`
//			IsFamilySafe         bool     `json:"isFamilySafe"`
//			AvailableCountries   []string `json:"availableCountries"`
//			IsUnlisted           bool     `json:"isUnlisted"`
//			HasYpcMetadata       bool     `json:"hasYpcMetadata"`
//			ViewCount            string   `json:"viewCount"`
//			Category             string   `json:"category"`
//			PublishDate          string   `json:"publishDate"`
//			OwnerChannelName     string   `json:"ownerChannelName"`
//			LiveBroadcastDetails struct {
//				IsLiveNow      bool      `json:"isLiveNow"`
//				StartTimestamp time.Time `json:"startTimestamp"`
//				EndTimestamp   time.Time `json:"endTimestamp"`
//			} `json:"liveBroadcastDetails"`
//			UploadDate string `json:"uploadDate"`
//		} `json:"playerMicroformatRenderer"`
//	} `json:"microformat"`
//	Cards struct {
//		CardCollectionRenderer struct {
//			Cards []struct {
//				CardRenderer struct {
//					Teaser struct {
//						SimpleCardTeaserRenderer struct {
//							Message struct {
//								SimpleText string `json:"simpleText"`
//							} `json:"message"`
//							TrackingParams       string `json:"trackingParams"`
//							Prominent            bool   `json:"prominent"`
//							LogVisibilityUpdates bool   `json:"logVisibilityUpdates"`
//							OnTapCommand         struct {
//								ClickTrackingParams                   string `json:"clickTrackingParams"`
//								ChangeEngagementPanelVisibilityAction struct {
//									TargetID   string `json:"targetId"`
//									Visibility string `json:"visibility"`
//								} `json:"changeEngagementPanelVisibilityAction"`
//							} `json:"onTapCommand"`
//						} `json:"simpleCardTeaserRenderer"`
//					} `json:"teaser"`
//					CueRanges []struct {
//						StartCardActiveMs string `json:"startCardActiveMs"`
//						EndCardActiveMs   string `json:"endCardActiveMs"`
//						TeaserDurationMs  string `json:"teaserDurationMs"`
//						IconAfterTeaserMs string `json:"iconAfterTeaserMs"`
//					} `json:"cueRanges"`
//					TrackingParams string `json:"trackingParams"`
//				} `json:"cardRenderer"`
//			} `json:"cards"`
//			HeaderText struct {
//				SimpleText string `json:"simpleText"`
//			} `json:"headerText"`
//			Icon struct {
//				InfoCardIconRenderer struct {
//					TrackingParams string `json:"trackingParams"`
//				} `json:"infoCardIconRenderer"`
//			} `json:"icon"`
//			CloseButton struct {
//				InfoCardIconRenderer struct {
//					TrackingParams string `json:"trackingParams"`
//				} `json:"infoCardIconRenderer"`
//			} `json:"closeButton"`
//			TrackingParams           string `json:"trackingParams"`
//			AllowTeaserDismiss       bool   `json:"allowTeaserDismiss"`
//			LogIconVisibilityUpdates bool   `json:"logIconVisibilityUpdates"`
//		} `json:"cardCollectionRenderer"`
//	} `json:"cards"`
//	TrackingParams string `json:"trackingParams"`
//	Attestation    struct {
//		PlayerAttestationRenderer struct {
//			Challenge    string `json:"challenge"`
//			BotguardData struct {
//				Program            string `json:"program"`
//				InterpreterSafeURL struct {
//					PrivateDoNotAccessOrElseTrustedResourceURLWrappedValue string `json:"privateDoNotAccessOrElseTrustedResourceUrlWrappedValue"`
//				} `json:"interpreterSafeUrl"`
//				ServerEnvironment int `json:"serverEnvironment"`
//			} `json:"botguardData"`
//		} `json:"playerAttestationRenderer"`
//	} `json:"attestation"`
//	AdBreakHeartbeatParams string `json:"adBreakHeartbeatParams"`
//	FrameworkUpdates       struct {
//		EntityBatchUpdate struct {
//			Mutations []struct {
//				EntityKey string `json:"entityKey"`
//				Type      string `json:"type"`
//				Payload   struct {
//					OfflineabilityEntity struct {
//						Key                     string `json:"key"`
//						AddToOfflineButtonState string `json:"addToOfflineButtonState"`
//					} `json:"offlineabilityEntity"`
//				} `json:"payload"`
//			} `json:"mutations"`
//			Timestamp struct {
//				Seconds string `json:"seconds"`
//				Nanos   int    `json:"nanos"`
//			} `json:"timestamp"`
//		} `json:"entityBatchUpdate"`
//	} `json:"frameworkUpdates"`
//}
