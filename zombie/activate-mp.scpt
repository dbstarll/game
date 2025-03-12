tell application "System Events"
  tell application process "WeChat"
    set frontmost to true
    click menu item "向僵尸开炮" of menu 1 of menu bar item "窗口" of menu bar 1
  end tell
end tell
