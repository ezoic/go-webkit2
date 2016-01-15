package webkit2

// #include <webkit2/webkit2.h>
import "C"
import "unsafe"
import "github.com/gotk3/gotk3/glib"

type Settings struct {
	*glib.Object
	settings *C.WebKitSettings
}

// newSettings creates a new Settings with default values.
//
// See also: webkit_settings_new at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitSettings.html#webkit-settings-new.
func newSettings(settings *C.WebKitSettings) *Settings {
	return &Settings{&glib.Object{glib.ToGObject(unsafe.Pointer(settings))}, settings}
}

// GetAutoLoadImages returns the "auto-load-images" property.
//
// See also: webkit_settings_get_auto_load_images at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitSettings.html#webkit-settings-get-auto-load-images
func (s *Settings) GetAutoLoadImages() bool {
	return gobool(C.webkit_settings_get_auto_load_images(s.settings))
}

// SetAutoLoadImages sets the "auto-load-images" property.
//
// See also: webkit_settings_get_auto_load_images at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitSettings.html#webkit-settings-set-auto-load-images
func (s *Settings) SetAutoLoadImages(autoLoad bool) {
	C.webkit_settings_set_auto_load_images(s.settings, gboolean(autoLoad))
}

// Disables majority of toggleable settings.
func (s *Settings) DisableToggleableSettings() {
	C.webkit_settings_set_auto_load_images(s.settings, gboolean(false))
	C.webkit_settings_set_enable_frame_flattening(s.settings, gboolean(false))
	C.webkit_settings_set_enable_html5_database(s.settings, gboolean(false))
	C.webkit_settings_set_enable_html5_local_storage(s.settings, gboolean(false))
	C.webkit_settings_set_enable_hyperlink_auditing(s.settings, gboolean(false))
	C.webkit_settings_set_enable_java(s.settings, gboolean(false))
	C.webkit_settings_set_enable_javascript(s.settings, gboolean(false))
	C.webkit_settings_set_enable_offline_web_application_cache(s.settings, gboolean(false))
	C.webkit_settings_set_enable_plugins(s.settings, gboolean(false))
	C.webkit_settings_set_enable_xss_auditor(s.settings, gboolean(false))
	C.webkit_settings_set_javascript_can_open_windows_automatically(s.settings, gboolean(false))
	C.webkit_settings_set_enable_private_browsing(s.settings, gboolean(false))
	C.webkit_settings_set_enable_developer_extras(s.settings, gboolean(false))
	C.webkit_settings_set_enable_resizable_text_areas(s.settings, gboolean(false))
	C.webkit_settings_set_enable_tabs_to_links(s.settings, gboolean(false))
	C.webkit_settings_set_enable_dns_prefetching(s.settings, gboolean(false))
	C.webkit_settings_set_enable_caret_browsing(s.settings, gboolean(false))
	C.webkit_settings_set_enable_fullscreen(s.settings, gboolean(false))
	C.webkit_settings_set_print_backgrounds(s.settings, gboolean(false))
	C.webkit_settings_set_enable_webaudio(s.settings, gboolean(false))
	C.webkit_settings_set_enable_webgl(s.settings, gboolean(false))
	C.webkit_settings_set_allow_modal_dialogs(s.settings, gboolean(false))
	C.webkit_settings_set_javascript_can_access_clipboard(s.settings, gboolean(false))
	C.webkit_settings_set_enable_page_cache(s.settings, gboolean(false))
	C.webkit_settings_set_enable_smooth_scrolling(s.settings, gboolean(false))
	C.webkit_settings_set_enable_accelerated_2d_canvas(s.settings, gboolean(false))
	C.webkit_settings_set_enable_media_stream(s.settings, gboolean(false))
	C.webkit_settings_set_enable_spatial_navigation(s.settings, gboolean(false))
	C.webkit_settings_set_enable_mediasource(s.settings, gboolean(false))
}

func (s *Settings) SetEnableJavascript(enableJavascript bool) {
	C.webkit_settings_set_enable_javascript(s.settings, gboolean(enableJavascript))
}

func (s *Settings) GetEnableJavascript() bool {
	return gobool(C.webkit_settings_get_enable_javascript(s.settings))
}

// SetUserAgentWithApplicationDetails sets the "user-agent" property by
// appending the application details to the default user agent.
//
// See also: webkit_settings_set_user_agent_with_application_details at
// http://webkitgtk.org/reference/webkit2gtk/unstable/WebKitSettings.html#webkit-settings-set-user-agent-with-application-details
func (s *Settings) SetUserAgentWithApplicationDetails(appName, appVersion string) {
	C.webkit_settings_set_user_agent_with_application_details(s.settings, (*C.gchar)(C.CString(appName)), (*C.gchar)(C.CString(appVersion)))
}
