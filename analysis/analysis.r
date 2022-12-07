library(ggplot2)
library(dplyr)
library(broom)
library(ggpubr)

deley_data <- read.table("./rawData/networkPropogationDeley.csv", header = TRUE, sep = ",")
extern_var <- read.table("./rawData/PercentageOfDeliveredMessages.csv", header = TRUE, sep = ",")

png(file = "./images/histExtVar.png", width = 960, height = 960)
hist(extern_var$extVar, col = "#0d500d")
dev.off()

deley_data_lm <- lm(CDP ~ Nodes, data = deley_data)

png(file = "./images/scatterPlot.png", width = 960, height = 960)
plot(CDP ~ Nodes, data = deley_data, col = "#125c12")
dev.off()

summary(deley_data_lm)

png(file = "./images/hist.png", width = 960, height = 960)
hist(deley_data$CDP, col = "#0d500d")
dev.off()

png(file = "./images/homoedecstisy.png", width = 960, height = 960)
par(mfrow = c(2, 2))
plot(deley_data_lm, col = "#4040c5")
par(mfrow = c(1, 1))
dev.off()

png(file = "images/graph.png", width = 960, height = 960)
deley_data_graph <- ggplot(deley_data, aes(x = Nodes, y = CDP)) + geom_point()

deley_data_graph <- deley_data_graph + geom_smooth(method = "lm")

deley_data_graph <- deley_data_graph +
  stat_regline_equation(label.x = 11.264, label.y = -100.839)


deley_data_graph +
  theme_bw() +
  labs(title = "CDP",
      x = "Nodes(10-105)",
      y = "messured Network delay of (1 to 1200)",
      col = "#125c12")

dev.off()